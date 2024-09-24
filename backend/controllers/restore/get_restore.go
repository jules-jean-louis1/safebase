package controllers

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func GetRestore(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	restoreService := services.NewRestoreService()

	restore, err := restoreService.GetRestoreByID(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, restore)
}
