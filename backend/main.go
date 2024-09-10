package main

import (
	databaseController "backend/controllers/database"
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

	router.DELETE("/delete-database/:id", func(c *gin.Context) {
		databaseController.DeleteDatabase(c)
	})

	router.Run("safebase:8080/api/v1")
}
