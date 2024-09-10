package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// Chargement des variables d'environnement
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found. Using environment variables.")
	}

	// Utilisation de DATABASE_URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Ajout de sslmode=disable si ce n'est pas déjà présent dans l'URL
	if !strings.Contains(dbURL, "sslmode=") {
		dbURL += "?sslmode=disable"
	}

	// Tentative de connexion à la base de données
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Si aucune erreur, on assigne la connexion à la variable globale DB
	DB = db
	fmt.Println("Successfully connected to the database")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}