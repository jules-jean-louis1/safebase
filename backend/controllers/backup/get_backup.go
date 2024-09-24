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

func GetFullBackups(c *gin.Context) {
	backupService := services.NewBackupService()

	backups, err := backupService.GetBackupsFull()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, backups)
}

func GetBackupOptions(c *gin.Context) {
	dbName := c.Query("dbName")
	dbType := c.Query("dbType")
	order := c.Query("order")

	backupService := services.NewBackupService()

	if dbName != "" {
		backups, err := backupService.GetBackupBy("database_name", dbName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, backups)
	}
	if dbType != "" {
		backups, err := backupService.GetBackupBy("type", dbType)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, backups)
	}
	if order != "" {
		backups, err := backupService.GetBackupBy("order", order)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, backups)
	}
}
