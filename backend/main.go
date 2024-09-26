package main

import (
	backupController "backend/controllers/backup"
	"backend/controllers/dashboard"
	databaseController "backend/controllers/database"
	"backend/controllers/execution"
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

	// Routes

	api := router.Group("/api")

	{

		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the Backup Service",
			})
		})

		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.POST("/database", func(c *gin.Context) {
			databaseController.AddDatabase(c, cronService)
		})

		api.PUT("/database", func(c *gin.Context) {
			databaseController.UpdateDatabase(c, cronService)
		})

		api.GET("/database/:id", func(c *gin.Context) {
			databaseController.GetDatabaseByID(c)
		})

		api.GET("/databases", func(c *gin.Context) {
			databaseController.GetAllDatabases(c)
		})

		api.GET("/databases/options", func(c *gin.Context) {
			databaseController.GetDatabaseOptions(c)
		})

		api.DELETE("/database/:id", func(c *gin.Context) {
			databaseController.DeleteDatabase(c)
		})

		// Test Connection to db
		api.GET("/database/test", func(c *gin.Context) {
			databaseController.TestConnection(c)
		})

		// backup Route
		api.POST("/backup", func(c *gin.Context) {
			backupController.AddBackup(c)
		})

		api.GET("/backups", func(c *gin.Context) {
			backupController.GetBackups(c)
		})

		api.GET("/backups/options", func(c *gin.Context) {
			backupController.GetBackupOptions(c)
		})

		api.GET("/backups/full", func(c *gin.Context) {
			backupController.GetFullBackups(c)
		})

		api.GET("/get-backup/:id", func(c *gin.Context) {
			backupController.GetBackupByID(c)
		})

		api.DELETE("/backup/:id", func(c *gin.Context) {
			backupController.DeleteBackup(c)
		})

		// Restore Route

		api.POST("/restore", func(c *gin.Context) {
			restoreController.NewRestore(c)
		})

		api.POST("/delete-restore", func(c *gin.Context) {
			restoreController.DeleteRestore(c)
		})

		// Executions Route

		api.GET("/executions", func(c *gin.Context) {
			execution.GetExecutions(c)
		})

		// Dashboard Route

		api.GET("/dashboard", func(c *gin.Context) {
			dashboard.DashboardData(c)
		})

		// test route

		api.GET("/testCo", func(c *gin.Context) {
			databaseController.TestF(c)
		})
	}

	router.Run(":8080")
}
