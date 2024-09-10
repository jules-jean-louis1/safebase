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
