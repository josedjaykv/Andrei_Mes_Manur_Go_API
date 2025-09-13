package models

import "time"

type ResistanceReport struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	DaemonID    *uint        `gorm:"index" json:"daemon_id,omitempty"`
	VictimID    *uint        `json:"victim_id,omitempty"`
	ReportType  ReportType   `gorm:"type:varchar(50);not null" json:"report_type"`
	Description string       `gorm:"type:text;not null" json:"description"`
	Severity    Severity     `gorm:"type:varchar(20);default:'LOW'" json:"severity"`
	Status      ReportStatus `gorm:"type:varchar(20);default:'PENDING';index" json:"status"`
	CreatedAt   time.Time    `json:"created_at"`

	// Relaciones
	Daemon *User   `gorm:"foreignKey:DaemonID;constraint:OnDelete:SET NULL" json:"daemon,omitempty"`
	Victim *Victim `gorm:"foreignKey:VictimID;constraint:OnDelete:CASCADE" json:"victim,omitempty"`
}

func (ResistanceReport) TableName() string {
	return "resistance_reports"
}
