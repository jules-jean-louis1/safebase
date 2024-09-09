package services

import (
	"backend/db"
	"backend/model"

	"gorm.io/gorm"
)

type DatabaseService struct {
	DB *gorm.DB
}

func NewDatabaseService() *DatabaseService {
	return &DatabaseService{
		DB: db.GetDB(),
	}
}

// CreateDatabase permet de créer une nouvelle entrée dans la table Database
func (s *DatabaseService) CreateDatabase(
	name string,
	dbType string,
	host string,
	port string,
	username string,
	password string,
	databaseName string,
	connectionString string, // Facultatif
	cronSchedule string, // Facultatif
) (*model.Database, error) {

	// Création de la nouvelle instance du modèle Database avec les paramètres passés
	database := &model.Database{
		Name:             name,
		Type:             dbType,
		Host:             host,
		Port:             port,
		Username:         username,
		Password:         password,
		DatabaseName:     databaseName,
		ConnectionString: connectionString, // Facultatif
		CronSchedule:     cronSchedule,     // Facultatif
	}

	// Création de l'entrée dans la base de données
	result := s.DB.Create(database)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retourne l'objet Database créé
	return database, nil
}

// GetDatabaseByID permet de récupérer une entrée dans la table Database en fonction de son ID
func (s *DatabaseService) GetDatabaseByID(id string) (*model.Database, error) {
	// Création d'une nouvelle instance du modèle Database
	database := &model.Database{}

	// Recherche de l'entrée dans la base de données
	result := s.DB.First(database, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retourne l'objet Database trouvé
	return database, nil
}
