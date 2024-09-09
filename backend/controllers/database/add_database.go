package controllers

import (
	"backend/model"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func AddDatabase(c *gin.Context) {
	var database *model.Database

	if err := c.ShouldBindJSON(&database); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	databaseService := services.NewDatabaseService()

	// Save the database in the database
	database, err := databaseService.CreateDatabase(
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

	c.JSON(201, database)
}
