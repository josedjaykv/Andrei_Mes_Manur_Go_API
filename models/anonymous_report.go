package models

import "time"

type AnonymousReport struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Title      string         `gorm:"size:255" json:"title,omitempty"`
	Content    string         `gorm:"type:text;not null" json:"content"`
	Category   ReportCategory `gorm:"type:varchar(50);default:'GENERAL'" json:"category"`
	IsVerified bool           `gorm:"default:false" json:"is_verified"`
	CreatedAt  time.Time      `json:"created_at"`
}

func (AnonymousReport) TableName() string {
	return "anonymous_reports"
}
