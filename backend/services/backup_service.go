package services

import (
	"backend/db"
	"backend/model"

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

	result := s.DB.Where("id = ?", id).First(backup)
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

	result := s.DB.Find(&backups)
	if result.Error != nil {
		return nil, result.Error
	}

	return backups, nil
}

func (s *BackupService) GetBackupsFull() ([]model.Backup, error) {
	backups := []model.Backup{}

	result := db.GetDB().Preload("Database").Find(&backups)
	if result.Error != nil {
		return nil, result.Error
	}

	return backups, nil
}
