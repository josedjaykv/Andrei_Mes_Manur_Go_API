package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;size:100;not null" json:"username"`
	Email        string `gorm:"uniqueIndex;size:255;not null" json:"email"`
	PasswordHash string `gorm:"size:255;not null" json:"-"`
	Role         Role   `gorm:"type:varchar(20);not null" json:"role"`
	IsActive     bool   `gorm:"default:true" json:"is_active"`

	// Relaciones
	CapturedVictims   []Victim           `gorm:"foreignKey:DaemonID" json:"captured_victims,omitempty"`
	ResistanceReports []ResistanceReport `gorm:"foreignKey:DaemonID" json:"resistance_reports,omitempty"`
	RewardsReceived   []RewardPunishment `gorm:"foreignKey:DaemonID" json:"rewards_received,omitempty"`
	RewardsAssigned   []RewardPunishment `gorm:"foreignKey:AssignedBy" json:"rewards_assigned,omitempty"`
	DaemonStats       *DaemonStats       `gorm:"foreignKey:DaemonID" json:"daemon_stats,omitempty"`
}

func (User) TableName() string {
	return "users"
}

// Hook para validar el rol
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Role != RoleAndrei && u.Role != RoleDaemon && u.Role != RoleNetworkAdmin {
		return gorm.ErrInvalidData
	}
	return nil
}
