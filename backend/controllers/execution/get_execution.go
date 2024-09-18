package execution

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func GetExecutions(c *gin.Context) {
	backupService := services.NewBackupService()

	executions, err := backupService.GetExecutions()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, executions)

}
