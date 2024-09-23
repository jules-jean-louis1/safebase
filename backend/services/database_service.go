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
	isCronActive bool, // Facultatif
	cronSchedule string, // Facultatif
) (*model.Database, error) {

	// Création de la nouvelle instance du modèle Database avec les paramètres passés
	database := &model.Database{
		Name:         name,
		Type:         dbType,
		Host:         host,
		Port:         port,
		Username:     username,
		Password:     password,
		DatabaseName: databaseName,
		IsCronActive: isCronActive, // Facultatif
		CronSchedule: cronSchedule, // Facultatif
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

// GetAllDatabases permet de récupérer toutes les entrées de la table Database
func (s *DatabaseService) GetAllDatabases() ([]model.Database, error) {
	// Création d'une nouvelle instance du modèle Database
	var databases []model.Database

	// Recherche de toutes les entrées dans la base de données
	result := s.DB.Find(&databases).Order("created_at desc")
	if result.Error != nil {
		return nil, result.Error
	}

	// Retourne les objets Database trouvés
	return databases, nil
}

// UpdateDatabase permet de mettre à jour une entrée dans la table Database
func (s *DatabaseService) UpdateDatabase(
	id string,
	name string,
	dbType string,
	host string,
	port string,
	username string,
	password string,
	databaseName string,
	isCronActive bool, // Facultatif
	cronSchedule string, // Facultatif
) (*model.Database, error) {
	// Création d'une nouvelle instance du modèle Database
	database := &model.Database{}

	// Recherche de l'entrée dans la base de données
	result := s.DB.First(database, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Mise à jour des champs de l'objet Database
	database.Name = name
	database.Type = dbType
	database.Host = host
	database.Port = port
	database.Username = username
	database.Password = password
	database.DatabaseName = databaseName
	database.IsCronActive = isCronActive // Facultatif
	database.CronSchedule = cronSchedule // Facultatif

	// Mise à jour de l'entrée dans la base de données
	result = s.DB.Save(database)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retourne l'objet Database mis à jour
	return database, nil
}

func (s *DatabaseService) GetDatabaseBy(column string, value string) (*model.Database, error) {
	// Création d'une nouvelle instance du modèle Database
	database := &model.Database{}

	// Recherche de l'entrée dans la base de données
	result := s.DB.First(database, column+" = ?", value)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retourne l'objet Database trouvé
	return database, nil
}

func (s *DatabaseService) DeleteDatabase(id string) error {
	// Création d'une nouvelle instance du modèle Database
	database := &model.Database{}

	// Recherche de l'entrée dans la base de données
	result := s.DB.First(database, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// Suppression de l'entrée dans la base de données
	result = s.DB.Delete(database)
	if result.Error != nil {
		return result.Error
	}

	// Aucune erreur, retourne nil
	return nil
}
