package controllers

import (
	"backend/model"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func UpdateDatabase(c *gin.Context) {
	var database *model.Database

	if err := c.ShouldBindJSON(&database); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	databaseService := services.NewDatabaseService()

	if database.ID.String() == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	databaseResult, err := databaseService.UpdateDatabase(
		database.ID.String(),
		database.Name,
		database.Type,
		database.Host,
		database.Port,
		database.Username,
		database.Password,
		database.DatabaseName,
		database.ConnectionString,
		database.CronSchedule,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, databaseResult)
}
