package test

import (
	"backend/db"
	"backend/model"
	"backend/services"
	"os"
	"testing"
)

func TestInsertDatabase(t *testing.T) {
	// Vérifier que les variables d'environnement sont présentes
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		t.Fatal("DATABASE_URL environment variable is not set")
	}

	// Étape 1 : Initialiser la connexion à la base de données
	err := db.Connect()
	if err != nil {
		t.Fatalf("Erreur lors de la connexion à Safebase : %v", err)
	}

	// Étape 2 : Créer une instance de base de données pour le test
	database := model.Database{
		Name:         "test_database_safebase",
		Type:         "postgres",
		Host:         "localhost",
		Port:         os.Getenv("DB_PORT"),
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		IsCronActive: false,
		CronSchedule: "",
	}

	// Log des informations de connexion pour le débogage
	t.Logf("Tentative de connexion avec:")
	t.Logf("Host: %s", database.Host)
	t.Logf("Port: %s", database.Port)
	t.Logf("Database: %s", database.DatabaseName)

	// Le reste du test reste identique
	databaseService := services.NewDatabaseService()

	insertedDatabase, err := databaseService.CreateDatabase(
		database.Name,
		database.Type,
		database.Host,
		database.Port,
		database.Username,
		database.Password,
		database.DatabaseName,
		database.IsCronActive,
		database.CronSchedule,
	)

	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de la base de données : %v", err)
	}
	t.Logf("Base de données insérée avec succès : %v", insertedDatabase.Name)

	result, err := databaseService.GetDatabaseBy("name", database.Name)
	if err != nil {
		t.Fatalf("Erreur lors de la récupération de la base de données : %v", err)
	}

	if result.Name != database.Name {
		t.Errorf("Nom de la base de données attendu : %v, mais obtenu : %v", database.Name, result.Name)
	} else {
		t.Logf("Test réussi : la base de données a été insérée et récupérée correctement.")
	}

	err = databaseService.DeleteDatabase(result.ID.String())
	if err != nil {
		t.Fatalf("Erreur lors de la suppression de la base de données : %v", err)
	} else {
		t.Logf("Base de données supprimée avec succès.")
	}
}
