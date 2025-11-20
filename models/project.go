package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:150;not null" json:"name"`
	Slug        string         `gorm:"size:150;unique;not null" json:"slug"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedBy   uint           `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Owner   User            `gorm:"foreignKey:CreatedBy" json:"owner,omitempty"`
	Members []ProjectMember `json:"members,omitempty"`
	Tickets []Ticket        `json:"tickets,omitempty"`
}

type ProjectMember struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProjectID uint      `gorm:"not null;index:idx_project_user,unique" json:"project_id"`
	UserID    uint      `gorm:"not null;index:idx_project_user,unique" json:"user_id"`
	Role      string    `gorm:"size:20;default:'member';not null" json:"role"`
	CreatedAt time.Time `json:"created_at"`

	User    User    `json:"user,omitempty"`
	Project Project `json:"project,omitempty"`
}

type ProjectInvite struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	ProjectID uint       `gorm:"not null" json:"project_id"`
	Code      string     `gorm:"size:64;index" json:"code"`
	Email     string     `gorm:"size:150" json:"email"`
	Type      string     `gorm:"size:20;default:'code';not null" json:"type"`
	CreatedBy uint       `json:"created_by"`
	MaxUses   int        `gorm:"default:1" json:"max_uses"`
	Uses      int        `gorm:"default:0" json:"uses"`
	ExpiresAt *time.Time `json:"expires_at"`
	Status    string     `gorm:"size:20;default:'active'" json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}
