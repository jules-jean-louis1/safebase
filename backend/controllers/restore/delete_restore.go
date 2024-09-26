package restore

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func DeleteRestore(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
	}

	restoreService := services.NewRestoreService()

	err := restoreService.DeleteRestore(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"succes": "Restore deleted successfully"})
}
