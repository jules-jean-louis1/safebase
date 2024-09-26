package backup

import (
	"backend/services"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DeleteBackup(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	backupService := services.NewBackupService()

	// Récupérer les informations de la sauvegarde
	backup, err := backupService.GetBackupByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Chemin du fichier de sauvegarde
	filePath := filepath.Join("/app/backups", backup.Filename)

	// Vérifier si le fichier existe et le supprimer s'il existe
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			c.JSON(500, gin.H{"error": "Error deleting backup file", "details": err.Error()})
			return
		}
	} else if !os.IsNotExist(err) {
		// Une erreur autre que "fichier n'existe pas" s'est produite
		c.JSON(500, gin.H{"error": "Error checking backup file", "details": err.Error()})
		return
	}

	// Supprimer l'enregistrement de la base de données
	if err := backupService.DeleteBackup(id); err != nil {
		c.JSON(500, gin.H{"error": "Error deleting backup record", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Backup deleted successfully"})
}
