package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	ProjectID     uint       `json:"project_id"`
	Title         string     `gorm:"size:255;not null" json:"title"`
	Description   string     `gorm:"type:text" json:"description"`
	Type          string     `gorm:"size:20;default:'task';not null" json:"type"`
	Priority      string     `gorm:"size:20;default:'medium';not null" json:"priority"`
	Status        string     `gorm:"size:30;default:'todo';not null" json:"status"`
	EstimateHours int        `json:"estimate_hours"`
	DueDate       *time.Time `json:"due_date"`

	CreatorID  uint  `json:"creator_id"`
	AssigneeID *uint `json:"assignee_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Project  Project         `json:"project,omitempty"`
	Creator  User            `gorm:"foreignKey:CreatorID" json:"creator,omitempty"`
	Assignee *User           `gorm:"foreignKey:AssigneeID" json:"assignee,omitempty"`
	Comments []Comment       `json:"comments,omitempty"`
	History  []TicketHistory `json:"history,omitempty"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TicketID  uint      `gorm:"not null" json:"ticket_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`

	User User `json:"user,omitempty"`
}

type TicketHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TicketID  uint      `gorm:"not null" json:"ticket_id"`
	UserID    *uint     `json:"user_id"`
	Action    string    `gorm:"size:100" json:"action"`
	OldValue  string    `gorm:"type:text" json:"old_value"`
	NewValue  string    `gorm:"type:text" json:"new_value"`
	CreatedAt time.Time `json:"created_at"`
}
