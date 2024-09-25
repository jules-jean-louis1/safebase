package test

import (
	"backend/db"
	"backend/model"
	"backend/services"
	"os"
	"testing"
)

// Need to init the database connection before running the tests and use localdatabase
func TestInsertDB(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://postgres:password@localhost:5434/safebase?sslmode=disable&TimeZone=Europe/Paris")
	err := db.Connect()
	if err != nil {
		t.Fatalf("Erreur lors de la connexion à la base de données: %v", err)
	}

	database := model.Database{
		Name:         "testing_is_good",
		Type:         "postgres",
		Host:         "postgres_db",
		Port:         "5432",
		Username:     "postgres",
		Password:     "password",
		DatabaseName: "test",
	}

	databaseService := services.NewDatabaseService()

	_, err = databaseService.CreateDatabase(
		database.Name,
		database.Type,
		database.Host,
		database.Port,
		database.Username,
		database.Password,
		database.DatabaseName,
		false,
		"",
	)

	if err != nil {
		t.Errorf("Error inserting database: %v", err)
	}

	result, err := databaseService.GetDatabaseBy("Name", database.Name)

	if err != nil {
		t.Errorf("Error fetching database: %v", err)
	}

	if result.Name != database.Name {
		t.Errorf("Expected %v but got %v", database.Name, result.Name)
	} else {
		t.Logf("Database inserted successfully")
	}

	databaseService.DeleteDatabase(result.ID.String())

}
