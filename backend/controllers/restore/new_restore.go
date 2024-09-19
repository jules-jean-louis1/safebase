package controllers

import (
	"backend/model"
	"backend/services"
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type RestoreParams struct {
	BackupId            string `json:"backup_id"`
	DatabaseIdToRestore string `json:"database_id"`
}

func NewRestore(c *gin.Context) {
	var restoreParams RestoreParams

	// Vérification du format du JSON
	if err := c.ShouldBindJSON(&restoreParams); err != nil {
		c.JSON(400, gin.H{"Wrong format": err.Error()})
		return
	}

	// Vérification des paramètres requis
	if restoreParams.BackupId == "" {
		c.JSON(400, gin.H{"Message": "backup_id is required"})
		return
	}

	if restoreParams.DatabaseIdToRestore == "" {
		c.JSON(400, gin.H{"Message": "database_id is required"})
		return
	}

	// Récupération de la base de données et du backup
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

	// Restaurer la base de données en fonction du type
	var restoreErr error
	if database.Type == "mysql" {
		restoreErr = restoreMySQLDatabase(backup, database)
	} else if database.Type == "postgres" {
		restoreErr = restorePostgresDatabase(filepath, *database)
	}

	// Si une erreur survient durant la restauration
	if restoreErr != nil {
		c.JSON(500, gin.H{"Message": "Error restoring database", "Error": restoreErr.Error()})
		return
	}

	// Enregistrer les détails de la restauration
	restoreService := services.NewRestoreService()
	_, err = restoreService.CreateRestore(database.ID.String(), backup.ID.String(), "success", "", "", "")
	if err != nil {
		c.JSON(500, gin.H{"Message": "Error saving restore record", "Error": err.Error()})
		return
	}

	// Si tout se passe bien
	c.JSON(200, gin.H{"Message": "Database restored successfully"})
}

func isBackupFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func restorePostgresDatabase(backupFilePath string, database model.Database) error {
	// Connexion à la base de données cible pour exécuter les commandes de suppression et restauration
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.Username, database.Password, database.DatabaseName))
	if err != nil {
		return fmt.Errorf("error connecting to target database: %v", err)
	}
	defer db.Close()

	// Fermer toutes les connexions existantes à la base de données cible sauf la connexion actuelle
	_, err = db.Exec(fmt.Sprintf(`
		SELECT pg_terminate_backend(pid) 
		FROM pg_stat_activity 
		WHERE datname = '%s' AND pid <> pg_backend_pid()
	`, database.DatabaseName))
	if err != nil {
		return fmt.Errorf("error terminating existing connections: %v", err)
	}

	// Drop le schéma public (qui contient toutes les tables et objets) et recrée-le
	_, err = db.Exec("DROP SCHEMA public CASCADE")
	if err != nil {
		return fmt.Errorf("error dropping schema: %v", err)
	}

	_, err = db.Exec("CREATE SCHEMA public")
	if err != nil {
		return fmt.Errorf("error recreating schema: %v", err)
	}

	// Copier le fichier de sauvegarde dans le conteneur Docker
	containerPath := "/tmp/" + filepath.Base(backupFilePath)
	copyCmd := exec.Command("docker", "cp", backupFilePath, "plateforme-safebase-postgres_db-1:"+containerPath)
	if err := copyCmd.Run(); err != nil {
		return fmt.Errorf("error copying backup file to container: %v", err)
	}

	// Exécuter la commande de restauration depuis le fichier de sauvegarde
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

	// Supprimer le fichier de sauvegarde du conteneur Docker
	cleanupCmd := exec.Command("docker", "exec", "plateforme-safebase-postgres_db-1", "rm", containerPath)
	if err := cleanupCmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to remove temporary file from container: %v\n", err)
	}

	return nil
}

func restoreMySQLDatabase(backup *model.Backup, database *model.Database) error {
	oldDatabaseName := backup.Database.DatabaseName
	newDatabaseName := database.DatabaseName

	// File paths
	backupFilePath := "/app/backups/" + backup.Filename
	modifiedBackupFilePath := backupFilePath + ".modified.sql"

	// Debug: Log file paths for debugging
	fmt.Printf("Original dump file: %s\n", backupFilePath)
	fmt.Printf("Modified dump file: %s\n", modifiedBackupFilePath)

	// Create the modified dump file
	if err := createModifiedDumpFile(backupFilePath, modifiedBackupFilePath, newDatabaseName, oldDatabaseName); err != nil {
		return fmt.Errorf("error modifying dump file: %v", err)
	}

	// Connexion à MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", database.Username, database.Password, database.Host, database.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to MySQL: %v", err)
	}
	defer db.Close()

	// Fermer toutes les connexions existantes à la base de données cible
	rows, err := db.Query(fmt.Sprintf("SELECT ID FROM INFORMATION_SCHEMA.PROCESSLIST WHERE DB = '%s'", newDatabaseName))
	if err != nil {
		return fmt.Errorf("error fetching connection IDs: %v", err)
	}
	defer rows.Close()

	// Boucle pour tuer chaque connexion
	var connID int
	for rows.Next() {
		if err := rows.Scan(&connID); err != nil {
			return fmt.Errorf("error scanning connection ID: %v", err)
		}
		_, err := db.Exec(fmt.Sprintf("KILL %d", connID))
		if err != nil {
			fmt.Printf("Warning: Failed to kill connection ID %d: %v\n", connID, err)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error closing connections: %v", err)
	}

	// Supprimer la base de données si elle existe
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", newDatabaseName))
	if err != nil {
		return fmt.Errorf("error dropping database: %v", err)
	}

	// Créer une nouvelle base de données vide
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", newDatabaseName))
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	// Copier le fichier de sauvegarde modifié dans le conteneur
	containerPath := "/tmp/" + filepath.Base(modifiedBackupFilePath)
	copyCmd := exec.Command("docker", "cp", modifiedBackupFilePath, "plateforme-safebase-mysql_db-1:"+containerPath)
	if err := copyCmd.Run(); err != nil {
		return fmt.Errorf("error copying backup file to container: %v", err)
	}

	// Exécuter la commande de restauration
	restoreCmd := exec.Command("sh", "-c",
		fmt.Sprintf("cat %s | docker exec -i plateforme-safebase-mysql_db-1 /usr/bin/mysql -u%s -p%s %s",
			containerPath,
			database.Username,
			database.Password,
			newDatabaseName))

	var restoreOut bytes.Buffer
	restoreCmd.Stdout = &restoreOut
	restoreCmd.Stderr = &restoreOut

	if err := restoreCmd.Run(); err != nil {
		return fmt.Errorf("error restoring MySQL database: %v, output: %s", err, restoreOut.String())
	}

	// Supprimer le fichier de sauvegarde du conteneur
	cleanupCmd := exec.Command("docker", "exec", "plateforme-safebase-mysql_db-1", "rm", containerPath)
	if err := cleanupCmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to remove temporary file from container: %v\n", err)
	}

	return nil
}

func createModifiedDumpFile(originalFilePath, modifiedFilePath, newDatabaseName, oldDatabaseName string) error {

	inputFile, err := os.Open(originalFilePath)
	if err != nil {
		return fmt.Errorf("error opening original dump file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(modifiedFilePath)
	if err != nil {
		return fmt.Errorf("error creating modified dump file: %v", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()

		// Log each line read from the original file (optional, can be removed if the file is too large)
		fmt.Printf("Original line: %s\n", line)

		// Remove lines containing the sandbox mode comment (ensure exact match or trim spaces)
		if strings.Contains(strings.TrimSpace(line), "/*!999999- enable the sandbox mode */") {
			fmt.Println("Skipping sandbox mode line")
			continue
		}

		// Replace the old database name with the new one
		modifiedLine := strings.ReplaceAll(line, oldDatabaseName, newDatabaseName)

		// Log the modified line
		fmt.Printf("Modified line: %s\n", modifiedLine)

		writer.WriteString(modifiedLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from original dump file: %v", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error writing to modified dump file: %v", err)
	}

	// Log the completion of the file modification process
	fmt.Println("Modified dump file created successfully")

	// Cat the content of the modified file to the console for verification
	fmt.Println("Content of the modified dump file:")
	modifiedFileContent, err := os.ReadFile(modifiedFilePath)
	if err != nil {
		return fmt.Errorf("error reading modified dump file: %v", err)
	}
	fmt.Println(string(modifiedFileContent))

	return nil
}
