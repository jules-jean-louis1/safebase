package controllers

import (
	"backend/model"
	"backend/services"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

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

	if database.Host == "localhost" {
		database.Host = "host.docker.internal"
	}

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

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.Username, database.Password, database.DatabaseName))
	if err != nil {
		return fmt.Errorf("error connecting to target database: %v", err)
	}
	defer db.Close()

	// Vérifier si le fichier de sauvegarde existe localement
	if !isBackupFileExists(backupFilePath) {
		return fmt.Errorf("backup file does not exist at path: %s", backupFilePath)
	}

	// Construction de la commande psql pour restaurer la base de données
	restoreCmd := exec.Command(
		"psql",
		"-h", database.Host, // Hôte (nom de service Docker, ex. "postgres_service")
		"-U", database.Username, // Utilisateur PostgreSQL
		"-d", database.DatabaseName, // Nom de la base de données
		"-f", backupFilePath, // Fichier de sauvegarde à restaurer
	)

	// Injecter le mot de passe dans l'environnement pour psql
	restoreCmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", database.Password))

	// Exécuter la commande et capturer la sortie
	output, err := restoreCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error restoring PostgreSQL database: %v, output: %s", err, string(output))
	}

	// Si tout se passe bien, retour sans erreur
	return nil
}

func restoreMySQLDatabase(backupFilePath string, database model.Database) error {
	// Vérifier si le fichier de sauvegarde existe localement
	if !isBackupFileExists(backupFilePath) {
		return fmt.Errorf("le fichier de sauvegarde n'existe pas au chemin: %s", backupFilePath)
	}

	// Construire le Data Source Name (DSN) pour la connexion MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", database.Username, database.Password, database.Host, database.Port)

	// Ouvrir une connexion à MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erreur de connexion à MySQL: %v", err)
	}
	defer db.Close()

	// Vérifier que la connexion fonctionne
	if err := db.Ping(); err != nil {
		return fmt.Errorf("erreur de ping à la base de données: %v", err)
	}

	// Fermer toutes les connexions existantes à la base de données cible
	_, err = db.Exec(fmt.Sprintf("KILL (SELECT GROUP_CONCAT(ID) FROM INFORMATION_SCHEMA.PROCESSLIST WHERE DB = '%s')", database.DatabaseName))
	if err != nil {
		log.Printf("Avertissement: Échec de la fermeture des connexions existantes: %v\n", err)
	}

	// Supprimer la base de données si elle existe déjà
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", database.DatabaseName))
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression de la base de données: %v", err)
	}

	// Créer une nouvelle base de données vide
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", database.DatabaseName))
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la base de données: %v", err)
	}

	// Sélectionner la base de données fraîchement créée
	_, err = db.Exec(fmt.Sprintf("USE `%s`", database.DatabaseName))
	if err != nil {
		return fmt.Errorf("erreur lors de la sélection de la base de données: %v", err)
	}

	// Lire le fichier de sauvegarde depuis le chemin local
	restoreCmd := exec.Command(
		"mysql",
		"-h", database.Host, // Hôte (nom du service Docker, ex. "mysql_service")
		"-u", database.Username, // Utilisateur MySQL
		fmt.Sprintf("-p%s", database.Password), // Mot de passe MySQL (formaté directement)
		"-D", database.DatabaseName,            // Nom de la base de données à restaurer
	)

	// Ouvrir le fichier de sauvegarde en lecture
	backupFile, err := os.Open(backupFilePath)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier de sauvegarde: %v", err)
	}
	defer backupFile.Close()

	// Rediriger le contenu du fichier vers la commande mysql
	restoreCmd.Stdin = backupFile

	// Exécuter la commande de restauration et capturer la sortie
	output, err := restoreCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erreur lors de la restauration de la base de données MySQL: %v, sortie: %s", err, string(output))
	}

	// Si tout se passe bien, retour sans erreur
	return nil
}
