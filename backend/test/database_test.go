package test

import (
	"backend/db"
	"backend/model"
	"backend/services"
	"fmt"
	"os"
	"testing"
)

func TestInsertDatabase(t *testing.T) {
	// Étape 0 : Charger le fichier .env (si nécessaire)
	os.Setenv("DATABASE_URL", "postgresql://postgres:password@localhost:5434/safebase?sslmode=disable&TimeZone=Europe/Paris")
	err := db.Connect()
	if err != nil {
		fmt.Println("Info: .env file not found, using environment variables if set.")
	}

	// Étape 1 : Initialiser la connexion à la base de données (Safebase)
	err = db.Connect()
	if err != nil {
		t.Fatalf("Erreur lors de la connexion à Safebase : %v", err)
	}

	// Étape 2 : Créer une instance de base de données pour le test
	database := model.Database{
		Name:         "test_database_safebase", // Nom de la base de données pour le test
		Type:         "postgres",               // Type de base de données (ici PostgreSQL, utilisé par Safebase)
		Host:         "localhost",              // L'hôte, par exemple en local
		Port:         "5434",                   // Port PostgreSQL utilisé par Safebase
		Username:     "postgres",               // Nom d'utilisateur pour Safebase
		Password:     "password",               // Mot de passe pour Safebase
		DatabaseName: "safebase",               // Nom de la base de données Safebase
		IsCronActive: false,                    // Cron désactivé pour ce test
		CronSchedule: "",                       // Pas de planification de Cron
	}

	// Étape 3 : Initialisation du service de base de données
	databaseService := services.NewDatabaseService()

	// Étape 4 : Insérer la base de données
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

	// Étape 5 : Vérification de l'insertion
	if err != nil {
		t.Fatalf("Erreur lors de l'insertion de la base de données : %v", err)
	}
	t.Logf("Base de données insérée avec succès : %v", insertedDatabase.Name)

	// Étape 6 : Récupérer la base de données par son nom
	result, err := databaseService.GetDatabaseBy("name", database.Name)
	if err != nil {
		t.Fatalf("Erreur lors de la récupération de la base de données : %v", err)
	}

	// Étape 7 : Comparer les résultats pour vérifier que l'insertion a fonctionné
	if result.Name != database.Name {
		t.Errorf("Nom de la base de données attendu : %v, mais obtenu : %v", database.Name, result.Name)
	} else {
		t.Logf("Test réussi : la base de données a été insérée et récupérée correctement.")
	}

	// Étape 8 : Supprimer la base de données de test après le test
	err = databaseService.DeleteDatabase(result.ID.String())
	if err != nil {
		t.Fatalf("Erreur lors de la suppression de la base de données : %v", err)
	} else {
		t.Logf("Base de données supprimée avec succès.")
	}
}
