package main

import (
	backupController "backend/controllers/backup"
	databaseController "backend/controllers/database"
	restoreController "backend/controllers/restore"
	"backend/db"
	service "backend/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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

	// Configurer le middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://frontend:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/database", func(c *gin.Context) {
		databaseController.AddDatabase(c, cronService)
	})

	router.PUT("/update-database", func(c *gin.Context) {
		databaseController.UpdateDatabase(c, cronService)
	})

	router.GET("/database/:id", func(c *gin.Context) {
		databaseController.GetDatabaseByID(c)
	})

	router.GET("/databases", func(c *gin.Context) {
		databaseController.GetAllDatabases(c)
	})

	router.DELETE("/delete-database/:id", func(c *gin.Context) {
		databaseController.DeleteDatabase(c)
	})

	// Test Connection to db
	router.GET("/database/test", func(c *gin.Context) {
		databaseController.TestConnection(c)
	})

	// backup Route
	router.POST("/backup", func(c *gin.Context) {
		backupController.AddBackup(c)
	})

	router.GET("/backups", func(c *gin.Context) {
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
