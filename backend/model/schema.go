package model

import (
	"time"

	"github.com/google/uuid"
)

// Définition des rôles d'utilisateur
// Définition des rôles d'utilisateur
type UserRole string

const (
	AdminRole       UserRole = "admin"
	RegularUserRole UserRole = "user"
)

// Définition du statut des backups
type BackupStatus string

const (
	PendingBackupStatus    BackupStatus = "pending"
	InProgressBackupStatus BackupStatus = "in_progress"
	SuccessBackupStatus    BackupStatus = "success"
	FailedBackupStatus     BackupStatus = "failed"
)

// Définition du type de backup
type BackupType string

const (
	ManualBackupType    BackupType = "manual"
	ScheduledBackupType BackupType = "scheduled"
)

// Définition du statut des restaurations
type RestoreStatus string

const (
	PendingRestoreStatus    RestoreStatus = "pending"
	InProgressRestoreStatus RestoreStatus = "in_progress"
	SuccessRestoreStatus    RestoreStatus = "success"
	FailedRestoreStatus     RestoreStatus = "failed"
)

// Modèle User
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	Role      UserRole  `gorm:"type:user_role;default:'user';not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Modèle Database
type Database struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name             string    `gorm:"not null"`
	Type             string    `gorm:"not null"`
	Host             string    `gorm:"not null"`
	Port             string    `gorm:"not null"`
	Username         string    `gorm:"not null"`
	Password         string    `gorm:"not null"`
	DatabaseName     string    `gorm:"not null"`
	ConnectionString string    `gorm:"type:text"`
	CronSchedule     string    `gorm:"type:text"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

// Modèle Backup
type Backup struct {
	ID         uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DatabaseID uuid.UUID    `gorm:"type:uuid;not null"`
	Status     BackupStatus `gorm:"type:backup_status;not null"`
	BackupType BackupType   `gorm:"type:backup_type;not null"`
	Filename   string       `gorm:"not null"`
	Size       string       `gorm:"type:text"`
	ErrorMsg   string       `gorm:"type:text"`
	Log        string       `gorm:"type:text"`
	CreatedAt  time.Time    `gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `gorm:"autoUpdateTime"`
}

// Modèle Restore
type Restore struct {
	ID         uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DatabaseID uuid.UUID     `gorm:"type:uuid;not null"`
	BackupID   *uuid.UUID    `gorm:"type:uuid"`
	Status     RestoreStatus `gorm:"type:restore_status;not null"`
	Filename   string        `gorm:"not null"`
	ErrorMsg   string        `gorm:"type:text"`
	Log        string        `gorm:"type:text"`
	CreatedAt  time.Time     `gorm:"autoCreateTime"`
	UpdatedAt  time.Time     `gorm:"autoUpdateTime"`
}
