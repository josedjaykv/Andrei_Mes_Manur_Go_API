// models/reward_punishment.go
package models

import "time"

type RewardPunishment struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	DaemonID    uint       `gorm:"not null;index" json:"daemon_id"`
	AssignedBy  uint       `gorm:"not null" json:"assigned_by"`
	Type        RewardType `gorm:"type:varchar(20);not null" json:"type"`
	Title       string     `gorm:"size:255;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description,omitempty"`
	Points      int        `gorm:"default:0" json:"points"`
	CreatedAt   time.Time  `json:"created_at"`

	// Relaciones
	Daemon   User `gorm:"foreignKey:DaemonID;constraint:OnDelete:CASCADE" json:"daemon,omitempty"`
	Assigner User `gorm:"foreignKey:AssignedBy" json:"assigner,omitempty"`
}

func (RewardPunishment) TableName() string {
	return "rewards_punishments"
}
