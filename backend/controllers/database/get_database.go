package controllers

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func GetDatabaseByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	databaseService := services.NewDatabaseService()

	database, err := databaseService.GetDatabaseByID(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, database)
}

func GetAllDatabases(c *gin.Context) {
	databaseService := services.NewDatabaseService()

	databases, err := databaseService.GetAllDatabases()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, databases)
}

func GetDatabaseOptions(c *gin.Context) {
	dbName := c.Query("dbName")
	dbType := c.Query("dbType")
	isCronActive := c.Query("is_cron_active")

	databaseService := services.NewDatabaseService()

	if dbName != "" {
		databases, err := databaseService.GetDatabaseBy("database_name", dbName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, databases)
	}
	if dbType != "" {
		databases, err := databaseService.GetDatabaseBy("type", dbType)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, databases)

	}
	if isCronActive != "" {
		databases, err := databaseService.GetDatabaseBy("is_cron_active", isCronActive)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, databases)
	}

}
