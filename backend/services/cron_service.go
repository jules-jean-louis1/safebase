package services

import (
	"backend/db"
	model "backend/model"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

type CronService struct {
	DB        *gorm.DB
	Scheduler gocron.Scheduler
}

func NewCronService() (*CronService, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &CronService{
		DB:        db.GetDB(),
		Scheduler: scheduler,
	}, nil
}

func (s *CronService) StartCronJobs() error {
	err := s.updateCronJobs()
	if err != nil {
		return err
	}
	s.Scheduler.Start()
	return nil
}

func (s *CronService) updateCronJobs() error {
	var databases []model.Database
	result := s.DB.Where("is_cron_active = ?", true).Find(&databases)
	if result.Error != nil {
		return result.Error
	}

	// s.Scheduler.Clear()

	for _, database := range databases {
		err := s.addCronJob(database)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CronService) addCronJob(db model.Database) error {
	_, err := s.Scheduler.NewJob(
		gocron.CronJob(db.CronSchedule, false),
		gocron.NewTask(
			func() {
				// Ajoutez ici la logique pour effectuer le backup
				// Par exemple :
				// backupService.PerformBackup(db)
			},
		),
	)
	return err
}

func (s *CronService) RefreshCronJobs() error {
	return s.updateCronJobs()
}

func (s *CronService) Stop() {
	s.Scheduler.Shutdown()
}
