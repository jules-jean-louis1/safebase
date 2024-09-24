package services

import (
	utils "backend/controllers/utils"
	"backend/model"
	"log"

	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ScheduleBackup(database model.Database) {
	backupService := NewBackupService()

	filename := fmt.Sprintf("%s_%s.sql", database.Name, time.Now().Format("2006-01-02-15:04:05"))
	directory := "/app/backups/"
	filepath := directory + filename
	var status string

	// Ensure the directory exists
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		status = "failed"
		log.Printf("Error creating backup directory: %v", err)

		// Log backup failure for directory creation error
		backup := model.Backup{
			DatabaseID: database.ID,
			Status:     model.BackupStatus(status),
			BackupType: "scheduled",
			Filename:   filename,
			Size:       "0",
			ErrorMsg:   "Error creating backup directory: " + err.Error(),
			Log:        "",
		}
		_, err := backupService.CreateBackup(
			backup.DatabaseID.String(),
			backup.Status,
			backup.BackupType,
			backup.Filename,
			backup.Size,
			backup.ErrorMsg,
			backup.Log,
		)
		if err != nil {
			log.Printf("Error creating backup record: %v", err)
		}
		return
	}

	// Perform the actual database backup
	if err := performDatabaseBackup(filepath, database); err != nil {
		status = "failed"
	} else {
		status = "success"
	}

	// Get the size of the backup file
	size, err := utils.GetSizeBackup(filepath)
	if err != nil {
		size = 0
	}

	// Create the backup object
	backup := model.Backup{
		DatabaseID: database.ID,
		Status:     model.BackupStatus(status),
		BackupType: "scheduled",
		Filename:   filename,
		Size:       fmt.Sprintf("%d", size),
		ErrorMsg: func() string {
			if err != nil {
				return err.Error()
			}
			return ""
		}(),
		Log: "",
	}

	// Insert the backup record into the database
	_, err = backupService.CreateBackup(
		backup.DatabaseID.String(),
		backup.Status,
		backup.BackupType,
		backup.Filename,
		backup.Size,
		backup.ErrorMsg,
		backup.Log,
	)

	// Log any error during the backup creation
	if err != nil {
		log.Printf("Error creating backup record: %v", err)
		os.Remove(filepath) // Optionally remove the backup file if DB insert fails
		return
	}

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
		cmd = exec.Command("mysqldump",
			"-h", database.Host,
			"-P", database.Port,
			"-u", database.Username,
			"-p"+database.Password,
			"--databases", database.DatabaseName)
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
