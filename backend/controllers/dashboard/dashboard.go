package dashboard

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func DashboardData(c *gin.Context) {

	DashboardService := services.NewDashboardService()

	data, err := DashboardService.GetDashboardData()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}
