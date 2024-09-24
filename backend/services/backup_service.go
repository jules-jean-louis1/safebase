package services

import (
	"backend/db"
	"backend/model"
	"sort"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BackupService struct {
	DB *gorm.DB
}

func NewBackupService() *BackupService {
	return &BackupService{
		DB: db.GetDB(),
	}
}

func (s *BackupService) CreateBackup(
	databaseID string,
	status model.BackupStatus,
	backupType model.BackupType,
	filename string,
	size string,
	errorMsg string,
	log string,
) (*model.Backup, error) {

	dbID, err := uuid.Parse(databaseID)
	if err != nil {
		return nil, err
	}

	backup := &model.Backup{
		DatabaseID: dbID,
		Status:     status,
		BackupType: backupType,
		Filename:   filename,
		Size:       size,
		ErrorMsg:   errorMsg,
		Log:        log,
	}

	result := s.DB.Create(backup)
	if result.Error != nil {
		return nil, result.Error
	}

	return backup, nil
}

func (s *BackupService) GetBackupByID(id string) (*model.Backup, error) {
	backup := &model.Backup{}

	result := s.DB.Preload("Database").Where("id = ?", id).First(backup)
	if result.Error != nil {
		return nil, result.Error
	}

	return backup, nil
}

func (s *BackupService) GetBackupsByDatabaseID(databaseID string) ([]model.Backup, error) {
	dbID, err := uuid.Parse(databaseID)
	if err != nil {
		return nil, err
	}

	backups := []model.Backup{}

	result := s.DB.Where("database_id = ?", dbID).Find(&backups)
	if result.Error != nil {
		return nil, result.Error
	}

	return backups, nil
}

func (s *BackupService) UpdateBackup(
	id string,
	status model.BackupStatus,
	backupType model.BackupType,
	filename string,
	size string,
	errorMsg string,
	log string,
) (*model.Backup, error) {
	backup := &model.Backup{}

	result := s.DB.First(backup, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	backup.Status = status
	backup.BackupType = backupType
	backup.Filename = filename
	backup.Size = size
	backup.ErrorMsg = errorMsg
	backup.Log = log

	result = s.DB.Save(backup)
	if result.Error != nil {
		return nil, result.Error
	}

	return backup, nil
}

func (s *BackupService) DeleteBackup(id string) error {
	backup := &model.Backup{}

	result := s.DB.First(backup, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = s.DB.Delete(backup)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *BackupService) GetBackups() ([]model.Backup, error) {
	backups := []model.Backup{}

	result := s.DB.Find(&backups).Order("created_at desc")
	if result.Error != nil {
		return nil, result.Error
	}

	return backups, nil
}

func (s *BackupService) GetBackupsFull() ([]model.Backup, error) {
	backups := []model.Backup{}

	result := db.GetDB().Preload("Database").Find(&backups).Order("created_at desc")
	if result.Error != nil {
		return nil, result.Error
	}

	return backups, nil
}

func (s *BackupService) GetBackupBy(field string, value string) ([]model.Backup, error) {
	backups := []model.Backup{}

	result := s.DB.Preload("Database").Joins("JOIN database ON database.id = backup.database_id").Where("database."+field+" = ?", value).Find(&backups)
	if result.Error != nil {
		return nil, result.Error
	}

	return backups, nil
}

func (s *BackupService) GetExecutions() (*model.Execution, error) {
	var backups []model.Backup
	var restores []model.Restore

	// Récupérer les backups
	result := s.DB.Preload("Database").Find(&backups)
	if result.Error != nil {
		return nil, result.Error
	}

	// Récupérer les restores
	result = s.DB.Preload("Database").Preload("Backup").Find(&restores)
	if result.Error != nil {
		return nil, result.Error
	}

	// Créer une liste d'éléments d'exécution
	var executionItems []model.ExecutionItem

	// Ajouter les backups à la liste
	for _, backup := range backups {
		var database model.Database
		if backup.Database != nil {
			database = *backup.Database // Déréférencement du pointeur
		}
		executionItems = append(executionItems, model.ExecutionItem{
			ID:         backup.ID.String(),
			Type:       "backup",
			Filename:   backup.Filename,
			Status:     string(model.BackupStatus(backup.Status)),
			Size:       backup.Size,
			DatabaseID: backup.DatabaseID.String(),
			Database:   database,
			CreatedAt:  backup.CreatedAt,
		})
	}

	// Ajouter les restores à la liste
	for _, restore := range restores {
		var database model.Database
		if restore.Database != nil {
			database = *restore.Database // Déréférencement du pointeur
		}
		executionItems = append(executionItems, model.ExecutionItem{
			ID:         restore.ID.String(),
			Type:       "restore",
			Filename:   restore.Backup.Filename, // Utiliser le nom de fichier du backup associé
			Status:     string(model.RestoreStatus(restore.Status)),
			Size:       restore.Backup.Size, // Utiliser la taille du backup associé
			DatabaseID: restore.DatabaseID.String(),
			Database:   database,
			CreatedAt:  restore.CreatedAt,
		})
	}

	// Trier les éléments par ordre décroissant de création
	sort.SliceStable(executionItems, func(i, j int) bool {
		return executionItems[i].CreatedAt.After(executionItems[j].CreatedAt)
	})

	// Retourner la structure Execution avec les éléments triés
	executions := &model.Execution{
		Items: executionItems,
	}

	return executions, nil
}
