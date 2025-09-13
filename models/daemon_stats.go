package models

import "time"

type DaemonStats struct {
	ID                         uint      `gorm:"primaryKey" json:"id"`
	DaemonID                   uint      `gorm:"uniqueIndex;not null" json:"daemon_id"`
	TotalCaptures              int       `gorm:"default:0" json:"total_captures"`
	SuccessfulConversions      int       `gorm:"default:0" json:"successful_conversions"`
	EscapesAllowed             int       `gorm:"default:0" json:"escapes_allowed"`
	ResistanceReportsSubmitted int       `gorm:"default:0" json:"resistance_reports_submitted"`
	TotalPoints                int       `gorm:"default:0;index:idx_points,sort:desc" json:"total_points"`
	RankPosition               *int      `json:"rank_position,omitempty"`
	UpdatedAt                  time.Time `json:"updated_at"`

	// Relaciones
	Daemon User `gorm:"foreignKey:DaemonID;constraint:OnDelete:CASCADE" json:"daemon,omitempty"`
}

func (DaemonStats) TableName() string {
	return "daemon_stats"
}
