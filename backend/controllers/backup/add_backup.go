package controllers

import (
	"backend/services"
	"fmt"
	"os"
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
}
