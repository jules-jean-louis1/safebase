package backup

import (
	"github.com/gin-gonic/gin"
)

func UploadBackup(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.SaveUploadedFile(file, "backups/"+file.Filename)

	c.JSON(200, gin.H{"file": file.Filename})
}
