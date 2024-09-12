package services

import (
	"backend/db"
	model "backend/model"
	"log"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CronService struct {
	DB        *gorm.DB
	Scheduler gocron.Scheduler
	Jobs      map[string]gocron.Job
}

func NewCronService() (*CronService, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &CronService{
		DB:        db.GetDB(),
		Scheduler: scheduler,
		Jobs:      make(map[string]gocron.Job),
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

	activeDBs := make(map[string]bool)
	for _, database := range databases {
		dbID := database.ID.String()
		activeDBs[dbID] = true
		if _, exists := s.Jobs[dbID]; !exists {
			err := s.addCronJob(database)
			if err != nil {
				return err
			}
		}
	}

	for dbID, job := range s.Jobs {
		if !activeDBs[dbID] {
			s.Scheduler.RemoveJob(job.ID())
			delete(s.Jobs, dbID)
		}
	}

	return nil
}

func (s *CronService) addCronJob(db model.Database) error {
	job, err := s.Scheduler.NewJob(
		gocron.CronJob(db.CronSchedule, false),
		gocron.NewTask(
			func() {
				log.Println("Running scheduled backup for database", db.ID)
				// ScheduleBackup(db)
			},
		),
	)
	if err != nil {
		return err
	}

	s.Jobs[db.ID.String()] = job
	return nil
}

func (s *CronService) RefreshCronJobs() error {
	return s.updateCronJobs()
}

func (s *CronService) Stop() {
	s.Scheduler.Shutdown()
}

func (s *CronService) AddOrUpdateJob(db model.Database) error {
	dbID := db.ID.String()
	if existingJob, exists := s.Jobs[dbID]; exists {
		s.Scheduler.RemoveJob(existingJob.ID())
	}
	return s.addCronJob(db)
}

func (s *CronService) RemoveJob(dbID uuid.UUID) error {
	strID := dbID.String()
	if job, exists := s.Jobs[strID]; exists {
		s.Scheduler.RemoveJob(job.ID())
		delete(s.Jobs, strID)
	}
	return nil
}
