package models

import (
	"time"

	"gorm.io/gorm"
)

type Victim struct {
	gorm.Model
	DaemonID    uint         `gorm:"not null;index" json:"daemon_id"`
	Name        string       `gorm:"size:255;not null" json:"name"`
	Email       string       `gorm:"size:255" json:"email,omitempty"`
	CaptureDate time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"capture_date"`
	Status      VictimStatus `gorm:"type:varchar(50);default:'CAPTURED'" json:"status"`
	Notes       string       `gorm:"type:text" json:"notes,omitempty"`

	// Relaciones
	Daemon            User               `gorm:"foreignKey:DaemonID;constraint:OnDelete:CASCADE" json:"daemon,omitempty"`
	ResistanceReports []ResistanceReport `gorm:"foreignKey:VictimID" json:"resistance_reports,omitempty"`
}

func (Victim) TableName() string {
	return "victims"
}
