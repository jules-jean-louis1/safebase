package main

import (
	backupController "backend/controllers/backup"
	CronController "backend/controllers/cron"
	databaseController "backend/controllers/database"
	restoreController "backend/controllers/restore"
	"backend/db"
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

	// Start the server
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/add-database", func(c *gin.Context) {
		databaseController.AddDatabase(c)
	})

	router.POST("/update-database", func(c *gin.Context) {
		databaseController.UpdateDatabase(c)
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

	// Cron routes
	router.GET("/debug/run-cron", CronController.RunCron)

	router.Run(":8080")
}

//  test
