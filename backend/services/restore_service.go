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
