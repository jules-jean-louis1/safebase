// package db

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect() error {
// 	// Chargement des variables d'environnement
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Warning: .env file not found. Using environment variables.")
// 	}

// 	// Utilisation de DATABASE_URL
// 	dbURL := os.Getenv("DATABASE_URL")
// 	if dbURL == "" {
// 		return fmt.Errorf("DATABASE_URL environment variable is not set")
// 	}

// 	// Ajout de sslmode=disable si ce n'est pas déjà présent dans l'URL
// 	if !strings.Contains(dbURL, "sslmode=") {
// 		dbURL += "?sslmode=disable"
// 	}

// 	// Ajout du fuseau horaire à l'URL de connexion
// 	if !strings.Contains(dbURL, "TimeZone=") {
// 		dbURL += "&TimeZone=Europe/Paris"
// 	}

// 	// Tentative de connexion à la base de données
// 	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to database: %w", err)
// 	}

// 	// Définir le fuseau horaire local pour l'application
// 	loc, err := time.LoadLocation("Europe/Paris")
// 	if err != nil {
// 		return fmt.Errorf("failed to load location: %w", err)
// 	}
// 	time.Local = loc

// 	// Si aucune erreur, on assigne la connexion à la variable globale DB
// 	DB = db
// 	fmt.Println("Successfully connected to the database")
// 	return nil
// }

// func GetDB() *gorm.DB {
// 	return DB
// }
// package db

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// // Connect initialise la connexion à la base de données en fonction du type de base de données défini dans l'environnement
// func Connect() error {
// 	// Chargement des variables d'environnement depuis le fichier .env
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Info: .env file not found, using environment variables if set.")
// 	}

// 	// Récupérer le type de base de données
// 	dbType := os.Getenv("DB_TYPE")
// 	fmt.Printf("Connecting to database type: %s\n", dbType)

// 	// Validation du type de base de données
// 	switch dbType {
// 	case "safebase":
// 		return connectToSafebase()
// 	case "postgres":
// 		return connectToPostgresProd()
// 	case "mysql":
// 		return connectToMySQL()
// 	default:
// 		return fmt.Errorf("unsupported database type: %s. Please set the 'DB_TYPE' environment variable to 'safebase', 'postgres', or 'mysql'", dbType)
// 	}
// }

// // Connexion à Safebase
// func connectToSafebase() error {
// 	fmt.Println("Connecting to Safebase...")
// 	dbUser := os.Getenv("DB_USER_SAFEBASE")
// 	dbPassword := os.Getenv("DB_PASSWORD_SAFEBASE")
// 	dbName := os.Getenv("DB_NAME_SAFEBASE")
// 	dbHost := os.Getenv("DB_HOST_SAFEBASE")
// 	dbPort := os.Getenv("DB_PORT_SAFEBASE")

// 	return connectPostgres(dbUser, dbPassword, dbHost, dbPort, dbName)
// }

// // Connexion à PostgreSQL Production
// func connectToPostgresProd() error {
// 	fmt.Println("Connecting to PostgreSQL Production...")
// 	dbUser := os.Getenv("DB_USER_POSTGRES")
// 	dbPassword := os.Getenv("DB_PASSWORD_POSTGRES")
// 	dbName := os.Getenv("DB_NAME_POSTGRES")
// 	dbHost := os.Getenv("DB_HOST_POSTGRES")
// 	dbPort := os.Getenv("DB_PORT_POSTGRES")

// 	return connectPostgres(dbUser, dbPassword, dbHost, dbPort, dbName)
// }

// // Connexion à MySQL
// func connectToMySQL() error {
// 	fmt.Println("Connecting to MySQL...")
// 	dbUser := os.Getenv("DB_USER_MYSQL")
// 	dbPassword := os.Getenv("MYSQL_PASSWORD")
// 	dbName := os.Getenv("DB_NAME_MYSQL")
// 	dbHost := os.Getenv("DB_HOST_MYSQL")
// 	dbPort := os.Getenv("DB_PORT_MYSQL")

// 	// Vérifier que les variables d'environnement sont bien définies
// 	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
// 		return fmt.Errorf("MySQL database environment variables are not properly set")
// 	}

// 	// Construire l'URL de connexion pour MySQL
// 	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

// 	// Tentative de connexion à la base de données MySQL
// 	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to MySQL database: %w", err)
// 	}

// 	DB = db
// 	fmt.Println("Successfully connected to the MySQL database")
// 	return nil
// }

// // Fonction pour connecter à PostgreSQL (commune à Safebase et Production)
// func connectPostgres(dbUser, dbPassword, dbHost, dbPort, dbName string) error {
// 	// Vérifier que les variables d'environnement sont bien définies
// 	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
// 		return fmt.Errorf("PostgreSQL database environment variables are not properly set")
// 	}

// 	// Construire l'URL de connexion pour PostgreSQL
// 	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

// 	// Tentative de connexion à la base de données PostgreSQL
// 	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to PostgreSQL database: %w", err)
// 	}

// 	// Définir le fuseau horaire local pour l'application
// 	loc, err := time.LoadLocation("Europe/Paris")
// 	if err != nil {
// 		return fmt.Errorf("failed to load location: %w", err)
// 	}
// 	time.Local = loc

// 	DB = db
// 	fmt.Println("Successfully connected to the PostgreSQL database")
// 	return nil
// }

// // GetDB renvoie l'instance de la base de données
// func GetDB() *gorm.DB {
// 	if DB == nil {
// 		fmt.Println("Database connection is not initialized. Call Connect() first.")
// 	}
// 	return DB
// }
package db

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initialise la connexion à la base de données Safebase uniquement
func Connect() error {
	// Chargement des variables d'environnement depuis le fichier .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Info: .env file not found, using environment variables if set.")
	}

	// Récupérer le type de base de données
	dbType := os.Getenv("DB_TYPE")
	if dbType != "safebase" {
		return fmt.Errorf("unsupported database type: %s. Only 'safebase' is supported", dbType)
	}

	// Connexion à Safebase
	dbUser := os.Getenv("DB_USER_SAFEBASE")
	dbPassword := os.Getenv("DB_PASSWORD_SAFEBASE")
	dbName := os.Getenv("DB_NAME_SAFEBASE")
	dbHost := os.Getenv("DB_HOST_SAFEBASE")
	dbPort := os.Getenv("DB_PORT_SAFEBASE")

	// Vérifier que les variables d'environnement sont bien définies
	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return fmt.Errorf("Safebase database environment variables are not properly set")
	}

	// Construire l'URL de connexion pour PostgreSQL (Safebase)
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Tentative de connexion à la base de données Safebase
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to Safebase database: %w", err)
	}

	// Définir le fuseau horaire local pour l'application
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return fmt.Errorf("failed to load location: %w", err)
	}
	time.Local = loc

	DB = db
	fmt.Println("Successfully connected to the Safebase database")
	return nil
}

// GetDB renvoie l'instance de la base de données
func GetDB() *gorm.DB {
	if DB == nil {
		fmt.Println("Database connection is not initialized. Call Connect() first.")
	}
	return DB
}
