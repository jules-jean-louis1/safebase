package services

import (
	utils "backend/controllers/utils"
	"backend/model"

	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ScheduleBackup(database model.Database) {
	// Code for scheduling backups
	var status string
	filename := fmt.Sprintf("%s_%s.sql", database.Name, time.Now().Format("2006-01-02-15:04:05"))
	directory := "/app/backups/"
	filepath := "/app/backups/" + filename

	// Ensure the directory exists
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		return
	}

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

	backup := model.Backup{
		DatabaseID: database.ID,
		Status:     model.BackupStatus(status),
		BackupType: "scheduled",
		Filename:   filename,
		Size:       fmt.Sprintf("%d", size),
		ErrorMsg:   err.Error(),
		Log:        "",
	}

	backupService := NewBackupService()

	_, err = backupService.CreateBackup(
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
		cmd = exec.Command("docker", "exec", "plateforme-safebase-mysql_db-1",
			"mysqldump",
			"-u", database.Username,
			"-p"+database.Password,
			"--databases", database.DatabaseName)
	} else if database.Type == "postgres" {
		cmd = exec.Command("docker", "exec", "plateforme-safebase-postgres_db-1",
			"pg_dump",
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
