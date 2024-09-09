package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// Construction de la chaîne de connexion à la base de données
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Tentative de connexion à la base de données
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Gestion de l'erreur de connexion (retourne l'erreur)
		return fmt.Errorf("échec de la connexion à la base de données: %w", err)
	}

	// Si aucune erreur, on assigne la connexion à la variable globale DB
	DB = db
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
