package services

import (
	"backend/db"
	"backend/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestoreService struct {
	DB *gorm.DB
}

func NewRestoreService() *RestoreService {
	return &RestoreService{
		DB: db.GetDB(),
	}
}

func (s *RestoreService) CreateRestore(
	databaseID string,
	backupID string,
	status model.RestoreStatus,
	filename string,
	errorMsg string,
	log string,
) (*model.Restore, error) {

	dbID, err := uuid.Parse(databaseID)
	if err != nil {
		return nil, err
	}

	bcID, err := uuid.Parse(backupID)
	if err != nil {
		return nil, err
	}

	restore := &model.Restore{
		DatabaseID: dbID,
		BackupID:   &bcID,
		Status:     status,
		Filename:   filename,
		ErrorMsg:   errorMsg,
		Log:        log,
	}

	result := s.DB.Create(restore)
	if result.Error != nil {
		return nil, result.Error
	}

	return restore, nil
}

func (s *RestoreService) GetRestoreByID(id string) (*model.Restore, error) {
	restore := &model.Restore{}

	result := s.DB.Where("id = ?", id).First(restore)
	if result.Error != nil {
		return nil, result.Error
	}

	return restore, nil
}

func (s *RestoreService) GetRestoresByDatabaseID(databaseID string) ([]model.Restore, error) {
	dbID, err := uuid.Parse(databaseID)
	if err != nil {
		return nil, err
	}

	var restores []model.Restore
	result := s.DB.Where("database_id = ?", dbID).Find(&restores)
	if result.Error != nil {
		return nil, result.Error
	}

	return restores, nil
}

func (s *RestoreService) GetRestoresByBackupID(backupID string) ([]model.Restore, error) {
	bcID, err := uuid.Parse(backupID)
	if err != nil {
		return nil, err
	}

	var restores []model.Restore
	result := s.DB.Where("backup_id = ?", bcID).Find(&restores)
	if result.Error != nil {
		return nil, result.Error
	}

	return restores, nil
}

func (s *RestoreService) UpdateRestore(
	id string,
	status model.RestoreStatus,
	errorMsg string,
	log string,
) (*model.Restore, error) {
	restore := &model.Restore{}

	result := s.DB.First(restore, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	restore.Status = status
	restore.ErrorMsg = errorMsg
	restore.Log = log

	result = s.DB.Save(restore)
	if result.Error != nil {
		return nil, result.Error
	}

	return restore, nil
}

func (s *RestoreService) DeleteRestore(id string) error {
	restore := &model.Restore{}

	result := s.DB.First(restore, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = s.DB.Delete(restore)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
