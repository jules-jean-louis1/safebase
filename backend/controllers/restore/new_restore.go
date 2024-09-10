package controllers

import (
	"backend/model"
	"backend/services"
	"database/sql"
	"fmt"
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
	} else if database.Type == "postgres" {
		err = restorePostgresDatabase(filepath, *database)
	}

	if err != nil {
		c.JSON(500, gin.H{"Message": "Error restoring database", "Error": err.Error()})
		return
	}

}

func isBackupFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func restoreMySQLDatabase(filepath string, database model.Database) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		database.Username, database.Password, database.Host, database.Port))
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Flush hosts
	_, err = db.Exec("FLUSH HOSTS")
	if err != nil {
		return fmt.Errorf("error flushing hosts: %v", err)
	}

	// Drop database
	_, err = db.Exec("DROP DATABASE IF EXISTS " + database.DatabaseName)
	if err != nil {
		return fmt.Errorf("error dropping database: %v", err)
	}

	// Create database
	_, err = db.Exec("CREATE DATABASE " + database.DatabaseName)
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	// Restore from file
	cmd := exec.Command("mysql", "-u", database.Username, "-p"+database.Password,
		"-h", database.Host, "-P", database.Port, database.DatabaseName)

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("error opening backup file: %v", err)
	}
	defer file.Close()

	cmd.Stdin = file
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error restoring database: %v, output: %s", err, string(output))
	}

	return nil
}

func restorePostgresDatabase(filepath string, database model.Database) error {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		database.Host, database.Port, database.Username, database.Password))
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Drop database
	_, err = db.Exec("DROP DATABASE IF EXISTS " + database.DatabaseName)
	if err != nil {
		return fmt.Errorf("error dropping database: %v", err)
	}

	// Create database
	_, err = db.Exec("CREATE DATABASE " + database.DatabaseName)
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	// Restore from file
	cmd := exec.Command("psql", "-U", database.Username, "-h", database.Host, "-p", database.Port, "-d", database.DatabaseName)

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("error opening backup file: %v", err)
	}
	defer file.Close()

	cmd.Stdin = file
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error restoring database: %v, output: %s", err, string(output))
	}

	return nil
}
