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

func (User) TableName() string {
	return "user"
}

// Modèle Database
type Database struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Type         string    `gorm:"not null" json:"type"`
	Host         string    `gorm:"not null" json:"host"`
	Port         string    `gorm:"not null" json:"port"`
	Username     string    `gorm:"not null" json:"username"`
	Password     string    `gorm:"not null" json:"password"`
	DatabaseName string    `gorm:"not null" json:"database_name"`
	IsCronActive bool      `gorm:"default:false" json:"is_cron_active"`
	CronSchedule string    `gorm:"type:text" json:"cron_schedule"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Database) TableName() string {
	return "database"
}

// Modèle Backup
type Backup struct {
	ID         uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DatabaseID uuid.UUID    `gorm:"type:uuid;not null"`
	Database   *Database    `gorm:"foreignKey:DatabaseID"`
	Status     BackupStatus `gorm:"type:backup_status;not null"`
	BackupType BackupType   `gorm:"type:backup_type;not null"`
	Filename   string       `gorm:"not null"`
	Size       string       `gorm:"type:text"`
	ErrorMsg   string       `gorm:"type:text"`
	Log        string       `gorm:"type:text"`
	CreatedAt  time.Time    `gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `gorm:"autoUpdateTime"`
}

func (Backup) TableName() string {
	return "backup"
}

// Modèle Restore
type Restore struct {
	ID         uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DatabaseID uuid.UUID     `gorm:"type:uuid;not null"`
	Database   *Database     `gorm:"foreignKey:DatabaseID"`
	BackupID   *uuid.UUID    `gorm:"type:uuid"`
	Backup     *Backup       `gorm:"foreignKey:BackupID"`
	Status     RestoreStatus `gorm:"type:restore_status;not null"`
	Filename   string        `gorm:"not null"`
	ErrorMsg   string        `gorm:"type:text"`
	Log        string        `gorm:"type:text"`
	CreatedAt  time.Time     `gorm:"autoCreateTime"`
	UpdatedAt  time.Time     `gorm:"autoUpdateTime"`
}

func (Restore) TableName() string {
	return "restore"
}
