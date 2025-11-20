package models

import (
	"encoding/json"
	"time"
)

type Badge struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Code        string    `gorm:"size:50;unique;not null" json:"code"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Icon        string    `gorm:"size:255" json:"icon"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserBadge struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `gorm:"not null;index:idx_user_badge,unique" json:"user_id"`
	BadgeID    uint      `gorm:"not null;index:idx_user_badge,unique" json:"badge_id"`
	AwardedFor string    `gorm:"size:255" json:"awarded_for"`
	CreatedAt  time.Time `json:"created_at"`

	Badge Badge `json:"badge,omitempty"`
}

type XPEvent struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Source    string    `gorm:"size:50;not null" json:"source"`
	SourceID  uint      `json:"source_id"`
	XPDelta   int       `gorm:"not null" json:"xp_delta"`
	CreatedAt time.Time `json:"created_at"`
}

type Challenge struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	ProjectID   *uint      `json:"project_id"`
	Title       string     `gorm:"size:150;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	XPReward    int        `gorm:"default:0" json:"xp_reward"`
	StartAt     *time.Time `json:"start_at"`
	EndAt       *time.Time `json:"end_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type ChallengeParticipant struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	ChallengeID uint            `gorm:"not null" json:"challenge_id"`
	UserID      uint            `gorm:"not null" json:"user_id"`
	Progress    json.RawMessage `gorm:"type:json" json:"progress"`
	Completed   bool            `gorm:"default:false" json:"completed"`
	CompletedAt *time.Time      `json:"completed_at"`
}

type LeaderboardSnapshot struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Period      string          `gorm:"size:20;not null" json:"period"`
	PeriodStart *time.Time      `gorm:"type:date" json:"period_start"`
	PeriodEnd   *time.Time      `gorm:"type:date" json:"period_end"`
	Snapshot    json.RawMessage `gorm:"type:json" json:"snapshot"`
	CreatedAt   time.Time       `json:"created_at"`
}
