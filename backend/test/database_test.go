package test

import (
	"backend/db"
	"backend/model"
	"backend/services"
	"os"
	"testing"
)

func TestInsertDatabase(t *testing.T) {
	// Étape 0 : Charger le fichier .env (si nécessaire)
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Fatal("DATABASE_URL environment variable is not set")
	}

	err := db.Connect()
	if err != nil {
		t.Fatalf("Erreur lors de la connexion à la base de données : %v", err)
	}

	// Étape 1 : Initialiser la connexion à la base de données (Safebase)
	err = db.Connect()
	if err != nil {
		t.Fatalf("Erreur lors de la connexion à Safebase : %v", err)
	}

	// Étape 2 : Créer une instance de base de données pour le test
	database := model.Database{
		Name:         "test_database_safebase",
		Type:         "postgres",
		Host:         "localhost",
		Port:         "5432",                   // Modifié pour correspondre au port du workflow
		Username:     os.Getenv("DB_USER"),     // Utiliser les variables d'environnement
		Password:     os.Getenv("DB_PASSWORD"), // du workflow
		DatabaseName: os.Getenv("DB_NAME"),
		IsCronActive: false,
		CronSchedule: "",
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
