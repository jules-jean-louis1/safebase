package controllers

import (
	utils "backend/controllers/utils"
	model "backend/model"
	"backend/services"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type BackupParams struct {
	DatabaseId string `json:"database_id"`
}

func AddBackup(c *gin.Context) {
	var backupParams *BackupParams

	if err := c.ShouldBindJSON(&backupParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if backupParams.DatabaseId == "" {
		c.JSON(400, gin.H{"error": "Missing required fields"})
		return
	}
	databaseService := services.NewDatabaseService()

	database, err := databaseService.GetDatabaseByID(backupParams.DatabaseId)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	filename := fmt.Sprintf("%s_%s.sql", database.Name, time.Now().Format("20060102150405"))
	directory := "../backups/"
	filepath := filepath.Join(directory, filename)

	// Ensure the directory exists
	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		c.JSON(500, gin.H{"Message": "Error creating backup directory", "Error": err.Error()})
		return
	}

	err = performDatabaseBackup(filepath, *database)
	if err != nil {
		c.JSON(500, gin.H{"Message": "Error creating backup", "Error": err.Error()})
		return
	}

	backup := model.Backup{
		DatabaseID: database.ID,
		Status:     "success",
		BackupType: "manual",
		Filename:   filename,
		Size:       "0",
		ErrorMsg:   "",
		Log:        "",
	}

	backupService := services.NewBackupService()

	backupResult, err := backupService.CreateBackup(backup.Filename, backup.Status, backup.BackupType, backup.Size, backup.ErrorMsg, backup.Log, backup.DatabaseID.String())

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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

	if database.Type == "mysql" {
		cmd = exec.Command("docker exec -it mysql_db mysqldump",
			"-u", database.Username,
			"-h", database.Host,
			database.DatabaseName)
		cmd.Env = append(os.Environ(), fmt.Sprintf("MYSQL_PWD=%s", database.Password))
	} else if database.Type == "postgres" {
		cmd = exec.Command("docker exec -it postgres_db pg_dump",
			"-U", database.Username,
			"-h", database.Host,
			"-p", database.Port,
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
