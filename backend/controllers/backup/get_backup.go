package controllers

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func GetBackups(c *gin.Context) {
	backupService := services.NewBackupService()

	backups, err := backupService.GetBackups()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, backups)
}

func GetBackupByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	backupService := services.NewBackupService()

	backup, err := backupService.GetBackupByID(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, backup)
}
