package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"size:100;not null" json:"name"`
	Username     string         `gorm:"size:50;unique;not null" json:"username"`
	Email        string         `gorm:"size:150;unique;not null" json:"email"`
	PasswordHash string         `gorm:"size:255;not null" json:"-"`
	Role         string         `gorm:"size:20;default:'developer';not null" json:"role"`
	TokenVersion int            `gorm:"default:0;not null" json:"token_version"`
	AvatarURL    string         `gorm:"size:255" json:"avatar_url"`
	LastActiveAt *time.Time     `json:"last_active_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	Stats         UserStats       `json:"stats,omitempty"`
	OwnedProjects []Project       `gorm:"foreignKey:CreatedBy" json:"owned_projects,omitempty"`
	Memberships   []ProjectMember `json:"memberships,omitempty"`
}

type UserStats struct {
	UserID             uint       `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	XP                 int64      `gorm:"default:0" json:"xp"`
	Level              int        `gorm:"default:1" json:"level"`
	TicketsClosedCount int        `gorm:"default:0" json:"tickets_closed_count"`
	LastActivity       *time.Time `json:"last_activity"`
}
