package services

import (
	"backend/db"

	"gorm.io/gorm"
)

type DashboardService struct {
	DB *gorm.DB
}

func NewDashboardService() *DashboardService {
	return &DashboardService{
		DB: db.GetDB(),
	}
}

type DatabaseData struct {
	Total    int
	Mysql    int
	Postgres int
}

type BackupData struct {
	Total      int
	Successful int
	Failed     int
}

type RestoreData struct {
	Total      int
	Successful int
	Failed     int
}

type DashboardData struct {
	Databases DatabaseData
	Backups   BackupData
	Restores  RestoreData
}

func (s *DashboardService) GetDashboardData() (*DashboardData, error) {
	var databases []struct {
		Total    int
		Mysql    int
		Postgres int
	}
	var backups []struct {
		Total      int
		Successful int
		Failed     int
	}
	var restores []struct {
		Total      int
		Successful int
		Failed     int
	}

	// Get database data
	s.DB.Raw(`
		SELECT
			COUNT(*) AS total,
			SUM(CASE WHEN type = 'mysql' THEN 1 ELSE 0 END) AS mysql,
			SUM(CASE WHEN type = 'postgres' THEN 1 ELSE 0 END) AS postgres
		FROM database
	`).Scan(&databases)

	// Get backup data
	s.DB.Raw(`
		SELECT
			COUNT(*) AS total,
			SUM(CASE WHEN status = 'success' THEN 1 ELSE 0 END) AS successful,
			SUM(CASE WHEN status = 'failed' THEN 1 ELSE 0 END) AS failed
		FROM backup
	`).Scan(&backups)

	// Get restore data
	s.DB.Raw(`
		SELECT
			COUNT(*) AS total,
			SUM(CASE WHEN status = 'success' THEN 1 ELSE 0 END) AS successful,
			SUM(CASE WHEN status = 'failed' THEN 1 ELSE 0 END) AS failed
		FROM restore
	`).Scan(&restores)

	return &DashboardData{
		Databases: DatabaseData{
			Total:    databases[0].Total,
			Mysql:    databases[0].Mysql,
			Postgres: databases[0].Postgres,
		},
		Backups: BackupData{
			Total:      backups[0].Total,
			Successful: backups[0].Successful,
			Failed:     backups[0].Failed,
		},
		Restores: RestoreData{
			Total:      restores[0].Total,
			Successful: restores[0].Successful,
			Failed:     restores[0].Failed,
		},
	}, nil
}
