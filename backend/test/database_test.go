package database_test

import (
	"backend/model"
	"backend/services"
	"testing"
)

func TestInsertDB(t *testing.T) {

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

	_, err := databaseService.CreateDatabase(
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
	}

}
