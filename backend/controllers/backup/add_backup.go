package controllers

import (
	utils "backend/controllers/utils"
	model "backend/model"
	"backend/services"

	"bytes"
	"fmt"

	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BackupParams struct {
	DatabaseId string `json:"database_id"`
}

func AddBackup(c *gin.Context) {
	backupParams := &BackupParams{}
	if err := c.ShouldBindJSON(&backupParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if backupParams.DatabaseId == "" {
		c.JSON(400, gin.H{"error": "Missing required fields"})
		return
	}
	backupParams.DatabaseId = strings.TrimSpace(backupParams.DatabaseId)

	// Vérifiez si l'UUID est valide
	if _, err := uuid.Parse(backupParams.DatabaseId); err != nil {
		fmt.Println("UUID invalide:", backupParams.DatabaseId)
		c.JSON(400, gin.H{"error": "Invalid UUID format"})
		return
	}
	databaseService := services.NewDatabaseService()

	database, err := databaseService.GetDatabaseByID(backupParams.DatabaseId)
	if err != nil {
		c.JSON(500, gin.H{"error_database": err.Error()})
		return
	}

	filename := fmt.Sprintf("%s_%s.sql", database.Name, time.Now().Format("2006-01-02-15:04:05"))
	directory := "/app/backups/"
	filepath := "/app/backups/" + filename

	// Ensure the directory exists
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": "Error creating backup directory", "details": err.Error()})
		return
	}

	if err := performDatabaseBackup(filepath, *database); err != nil {
		c.JSON(500, gin.H{"error": "Error creating backup", "details": err.Error()})
		return
	}

	// Get the size of the backup file
	size, err := GetSizeBackup(filepath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error verifying backup file", "details": err.Error()})
		return
	}
	backup := model.Backup{
		DatabaseID: database.ID,
		Status:     "success",
		BackupType: "manual",
		Filename:   filename,
		Size:       fmt.Sprintf("%d", size),
		ErrorMsg:   "",
		Log:        "",
	}

	backupService := services.NewBackupService()

	backupResult, err := backupService.CreateBackup(
		backup.DatabaseID.String(),
		backup.Status,
		backup.BackupType,
		backup.Filename,
		backup.Size,
		backup.ErrorMsg,
		backup.Log,
	)

	if err != nil {
		os.Remove(filepath)
		c.JSON(500, gin.H{"error": "Error saving backup record", "details": err.Error()})
		return
	}

	c.JSON(201, backupResult)
}

func performDatabaseBackup(filepath string, database model.Database) error {
	var cmd *exec.Cmd

	params := &utils.DBParams{
		Host:     database.Host,
		Port:     database.Port,
		Username: database.Username,
		Password: database.Password,
		DBName:   database.DatabaseName,
		DBType:   database.Type,
		SSLMode:  "disable",
	}

	co, err := utils.ConnectionTester(params)
	if err != nil {
		fmt.Println("Erreur de connexion:", err)
		return err
	}

	if !co {
		return fmt.Errorf("connection to database failed")
	}
	// TODO : tester la connexion à la base de données
	if database.Type == "mysql" {
		cmd = exec.Command("mysqldump",
			"--skip-comments",
			"-h", database.Host,
			"--port", database.Port,
			"--user", database.Username,
			"--password="+database.Password, database.DatabaseName)
	} else if database.Type == "postgres" {
		cmd = exec.Command("pg_dump",
			"-h", database.Host,
			"-U", database.Username,
			"--no-owner",
			"--no-acl",
			database.DatabaseName)
		cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", database.Password))
	} else {
		return fmt.Errorf("unsupported database type: %s", database.Type)
	}

	outfile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outfile.Close()

	cmd.Stdout = outfile

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing dump command: %v, stderr: %s", err, stderr.String())
	}

	fmt.Println("Backup created successfully at", filepath)
	return nil
}

func GetSizeBackup(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("error getting file info: %w", err)
	}
	return fi.Size(), nil
}
