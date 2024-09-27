package database

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func DeleteDatabase(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	databaseService := services.NewDatabaseService()

	err := databaseService.DeleteDatabase(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
