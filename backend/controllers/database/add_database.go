package controllers

import (
	utils "backend/controllers/utils"
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

	if database.Name == "" ||
		database.Type == "" ||
		database.Host == "" ||
		database.Port == "" ||
		database.Username == "" ||
		database.Password == "" ||
		database.DatabaseName == "" {
		c.JSON(400, gin.H{"error": "Missing required fields"})
		return
	}

	// Test the connection
	params := &utils.DBParams{
		Host:     database.Host,
		Port:     database.Port,
		Username: database.Username,
		Password: database.Password,
		DBName:   database.DatabaseName,
		SSLMode:  "disable",
		DBType:   database.Type,
	}

	_, err := utils.ConnectionTester(params)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	databaseService := services.NewDatabaseService()

	// Save the database in the database
	database, err = databaseService.CreateDatabase(
		database.Name,
		database.Type,
		database.Host,
		database.Port,
		database.Username,
		database.Password,
		database.DatabaseName,
		database.IsCronActive,
		database.CronSchedule,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, database)
}
