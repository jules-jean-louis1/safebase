package main

import (
	backupController "backend/controllers/backup"
	databaseController "backend/controllers/database"
	restoreController "backend/controllers/restore"
	"backend/db"
	service "backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connexion à la base de données avec gestion d'erreur
	err := db.Connect()
	if err != nil {
		// Affiche l'erreur et arrête le programme en cas d'échec
		log.Fatalf("Erreur lors de la connexion à la base de données: %v", err)
	}

	// Connexion réussie
	log.Println("Connexion à la base de données réussie")

	// Initialiser le service Cron
	cronService, err := service.NewCronService()
	if err != nil {
		log.Fatalf("Failed to initialize CronService: %v", err)
	}

	// Démarrer les tâches Cron
	err = cronService.StartCronJobs()
	if err != nil {
		log.Fatalf("Failed to start Cron jobs: %v", err)
	}

	// Start the server
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/add-database", func(c *gin.Context) {
		databaseController.AddDatabase(c, cronService)
	})

	router.PUT("/update-database", func(c *gin.Context) {
		databaseController.UpdateDatabase(c, cronService)
	})

	router.GET("/get-database/:id", func(c *gin.Context) {
		databaseController.GetDatabaseByID(c)
	})

	router.GET("/get-all-databases", func(c *gin.Context) {
		databaseController.GetAllDatabases(c)
	})

	router.DELETE("/delete-database/:id", func(c *gin.Context) {
		databaseController.DeleteDatabase(c)
	})

	// Test Connection to db
	router.GET("/test-connection", func(c *gin.Context) {
		databaseController.TestConnection(c)
	})

	// backup Route
	router.POST("/create-manual-backup", func(c *gin.Context) {
		backupController.AddBackup(c)
	})

	router.GET("/get-backups", func(c *gin.Context) {
		backupController.GetBackups(c)
	})

	router.GET("/get-backup/:id", func(c *gin.Context) {
		backupController.GetBackupByID(c)
	})

	router.DELETE("/delete-backup/:id", func(c *gin.Context) {
		backupController.DeleteBackup(c)
	})

	// Restore Route

	router.POST("/restore-database", func(c *gin.Context) {
		restoreController.NewRestore(c)
	})

	router.POST("/delete-restore", func(c *gin.Context) {
		restoreController.DeleteRestore(c)
	})

	router.Run(":8080")
}
