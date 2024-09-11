package controllers

import (
	"backend/model"
	"backend/services"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type RestoreParams struct {
	BackupId            string `json:"backup_id"`
	DatabaseIdToRestore string `json:"database_id"`
}

func NewRestore(c *gin.Context) {
	var restoreParams RestoreParams

	if err := c.ShouldBindJSON(&restoreParams); err != nil {
		c.JSON(400, gin.H{"Wrong format": err.Error()})
	}

	if restoreParams.BackupId == "" {
		c.JSON(400, gin.H{"Message": "backup_id is required"})
		return
	}

	if restoreParams.DatabaseIdToRestore == "" {
		c.JSON(400, gin.H{"Message": "database_id is required"})
		return
	}

	databaseService := services.NewDatabaseService()

	database, err := databaseService.GetDatabaseByID(restoreParams.DatabaseIdToRestore)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	backupService := services.NewBackupService()

	backup, err := backupService.GetBackupByID(restoreParams.BackupId)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	filepath := "/app/backups/" + backup.Filename

	if !isBackupFileExists(filepath) {
		c.JSON(400, gin.H{"Message": "Backup file does not exist"})
		return
	}

	if database.Type == "mysql" {
		err = restoreMySQLDatabase(filepath, *database)
		if err != nil {
			c.JSON(500, gin.H{"Message": "Error restoring database", "Error": err.Error()})
			return
		}
	} else if database.Type == "postgres" {
		err = restorePostgresDatabase(filepath, *database)
		if err != nil {
			c.JSON(500, gin.H{"Message": "Error restoring database", "Error": err.Error()})
			return
		}
	}

	if err != nil {
		c.JSON(500, gin.H{"Message": "Error restoring database", "Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Database restored successfully"})

}

func isBackupFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

// Donc, j'ai selectionné un backup et je selectionne une database à restaurer
// Copier le fichier dump.sql dans le conteneur
// docker cp dump.sql plateforme-safebase-postgres_db-1:/tmp/dump.sql
// Restaurer la base de données
// docker exec plateforme-safebase-postgres_db-1 psql -U postgres -d postgres -f /tmp/dump.sql
// Supprimer le fichier dump.sql du conteneur
// docker exec plateforme-safebase-postgres_db-1 rm /tmp/dump.sql
func restorePostgresDatabase(backupFilePath string, database model.Database) error {
	// Connexion à la base de données 'postgres' pour les opérations administratives
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		database.Host, database.Port, database.Username, database.Password))
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Fermer toutes les connexions existantes à la base de données cible
	_, err = db.Exec(fmt.Sprintf(`
		SELECT pg_terminate_backend(pid) 
		FROM pg_stat_activity 
		WHERE datname = '%s' AND pid <> pg_backend_pid()
	`, database.DatabaseName))
	if err != nil {
		return fmt.Errorf("error terminating existing connections: %v", err)
	}

	// Drop database si elle existe
	_, err = db.Exec("DROP DATABASE IF EXISTS " + database.DatabaseName)
	if err != nil {
		return fmt.Errorf("error dropping database: %v", err)
	}

	// Créer la nouvelle base de données
	_, err = db.Exec("CREATE DATABASE " + database.DatabaseName)
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	// Copier le fichier de sauvegarde dans le conteneur
	containerPath := "/tmp/" + filepath.Base(backupFilePath)
	copyCmd := exec.Command("docker", "cp", backupFilePath, "plateforme-safebase-postgres_db-1:"+containerPath)
	if err := copyCmd.Run(); err != nil {
		return fmt.Errorf("error copying backup file to container: %v", err)
	}

	// Exécuter la commande de restauration
	restoreCmd := exec.Command("docker", "exec", "plateforme-safebase-postgres_db-1",
		"psql",
		"-U", database.Username,
		"-d", database.DatabaseName,
		"-f", containerPath)

	restoreCmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", database.Password))
	output, err := restoreCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error restoring PostgreSQL database: %v, output: %s", err, string(output))
	}

	// Supprimer le fichier de sauvegarde du conteneur
	cleanupCmd := exec.Command("docker", "exec", "plateforme-safebase-postgres_db-1", "rm", containerPath)
	if err := cleanupCmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to remove temporary file from container: %v\n", err)
	}

	return nil
}

func restoreMySQLDatabase(backupFilePath string, database model.Database) error {
	// Connexion à MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		database.Username, database.Password, database.Host, database.Port))
	if err != nil {
		return fmt.Errorf("error connecting to MySQL: %v", err)
	}
	defer db.Close()

	// Fermer toutes les connexions existantes à la base de données cible
	_, err = db.Exec(fmt.Sprintf("KILL (SELECT GROUP_CONCAT(ID) FROM INFORMATION_SCHEMA.PROCESSLIST WHERE DB = '%s')", database.DatabaseName))
	if err != nil {
		fmt.Printf("Warning: Failed to kill existing connections: %v\n", err)
	}

	// Supprimer la base de données si elle existe
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", database.DatabaseName))
	if err != nil {
		return fmt.Errorf("error dropping database: %v", err)
	}

	// Créer une nouvelle base de données vide
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", database.DatabaseName))
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	// Copier le fichier de sauvegarde dans le conteneur
	containerPath := "/tmp/" + filepath.Base(backupFilePath)
	copyCmd := exec.Command("docker", "cp", backupFilePath, "plateforme-safebase-mysql_db-1:"+containerPath)
	if err := copyCmd.Run(); err != nil {
		return fmt.Errorf("error copying backup file to container: %v", err)
	}

	// Exécuter la commande de restauration
	restoreCmd := exec.Command("docker", "exec", "plateforme-safebase-mysql_db-1",
		"mysql",
		"-u", database.Username,
		fmt.Sprintf("-p%s", database.Password),
		database.DatabaseName,
		"-e", fmt.Sprintf("SOURCE %s", containerPath))

	output, err := restoreCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error restoring MySQL database: %v, output: %s", err, string(output))
	}

	// Supprimer le fichier de sauvegarde du conteneur
	cleanupCmd := exec.Command("docker", "exec", "plateforme-safebase-mysql_db-1", "rm", containerPath)
	if err := cleanupCmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to remove temporary file from container: %v\n", err)
	}

	return nil
}
