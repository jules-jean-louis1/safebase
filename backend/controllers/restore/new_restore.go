package controllers

import (
	"backend/model"
	"backend/services"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
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
	var errMsg string

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
		errMsg = fmt.Sprintf("Error retrieving database: %v", err)
		saveRestoreRecord(restoreParams, "failed", errMsg)
		c.JSON(500, gin.H{"error": errMsg})
		return
	}

	backupService := services.NewBackupService()
	backup, err := backupService.GetBackupByID(restoreParams.BackupId)
	if err != nil {
		errMsg = fmt.Sprintf("Error retrieving backup: %v", err)
		saveRestoreRecord(restoreParams, "failed", errMsg)
		c.JSON(500, gin.H{"error": errMsg})
		return
	}

	filepath := "/app/backups/" + backup.Filename
	if !isBackupFileExists(filepath) {
		errMsg = "Backup file does not exist"
		saveRestoreRecord(restoreParams, "failed", errMsg)
		c.JSON(400, gin.H{"Message": errMsg})
		return
	}

	// Restaurer la base de données en fonction du type
	var restoreErr error
	if database.Type == "mysql" {
		restoreErr = restoreMySQLDatabase(filepath, *database)
	} else if database.Type == "postgres" {
		restoreErr = restorePostgresDatabase(filepath, *database)
	}

	// Si une erreur survient durant la restauration
	if restoreErr != nil {
		errMsg = fmt.Sprintf("Error restoring database: %v", restoreErr)
		saveRestoreRecord(restoreParams, "failed", errMsg)
		c.JSON(500, gin.H{"Message": errMsg})
		return
	}

	// Enregistrer les détails de la restauration
	saveRestoreRecord(restoreParams, "success", "")

	// Si tout se passe bien
	c.JSON(200, gin.H{"Message": "Database restored successfully"})
}

func saveRestoreRecord(params RestoreParams, status string, errMsg string) {
	restoreService := services.NewRestoreService()
	_, err := restoreService.CreateRestore(params.DatabaseIdToRestore, params.BackupId, model.RestoreStatus(status), "", errMsg, "")
	if err != nil {
		fmt.Printf("Error saving restore record: %v\n", err)
	}
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

func restoreMySQLDatabase(backupFilePath string, database model.Database) error {
	// Connexion à MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		database.Username, database.Password, database.Host, database.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erreur de connexion à MySQL: %v", err)
	}
	defer db.Close()

	// Vérifier la connexion
	if err := db.Ping(); err != nil {
		return fmt.Errorf("erreur de ping à la base de données: %v", err)
	}

	// Fermer toutes les connexions existantes à la base de données cible
	_, err = db.Exec(fmt.Sprintf("KILL (SELECT GROUP_CONCAT(ID) FROM INFORMATION_SCHEMA.PROCESSLIST WHERE DB = '%s')", database.DatabaseName))
	if err != nil {
		log.Printf("Avertissement: Échec de la fermeture des connexions existantes: %v\n", err)
	}

	// Supprimer la base de données si elle existe
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", database.DatabaseName))
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression de la base de données: %v", err)
	}

	// Créer une nouvelle base de données vide
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", database.DatabaseName))
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la base de données: %v", err)
	}

	// Sélectionner la base de données
	_, err = db.Exec(fmt.Sprintf("USE `%s`", database.DatabaseName))
	if err != nil {
		return fmt.Errorf("erreur lors de la sélection de la base de données: %v", err)
	}

	// Lire le contenu du fichier de sauvegarde
	backupContent, err := ioutil.ReadFile(backupFilePath)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier de sauvegarde: %v", err)
	}

	// Diviser le contenu en instructions SQL individuelles
	statements := strings.Split(string(backupContent), ";")

	// Exécuter chaque instruction SQL individuellement
	for _, stmt := range statements {
		trimmedStmt := strings.TrimSpace(stmt)
		if trimmedStmt != "" {
			_, err := db.Exec(trimmedStmt)
			if err != nil {
				log.Printf("Avertissement: Erreur lors de l'exécution de l'instruction SQL: %v\nInstruction: %s\n", err, trimmedStmt)
				// Continuer l'exécution malgré l'erreur
			}
		}
	}

	return nil
}
