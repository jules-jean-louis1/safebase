package controllers

import (
	utils "backend/controllers/utils"
	"backend/model"
	services "backend/services"
	"fmt"
	"os"
	"time"
)

func ScheduleBackup(database model.Database, backup model.Backup) {
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

	backup = model.Backup{
		DatabaseID: database.ID,
		Status:     model.BackupStatus(status),
		BackupType: "scheduled",
		Filename:   filename,
		Size:       fmt.Sprintf("%d", size),
		ErrorMsg:   err.Error(),
		Log:        "",
	}

	backupService := services.NewBackupService()

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
