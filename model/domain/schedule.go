package domain

import "time"

type Schedule struct {
	ID           int      `gorm:"primaryKey"`
	MovieID      int      `gorm:"column:movies_id"`
	CinemaID     int      `gorm:"column:cinema_id"`
	StartTime    time.Time `gorm:"column:start_time"`
	ScreenNumber int       `gorm:"column:screen_number"`
	IsCancel     bool      `gorm:"column:is_cancle"` // sesuai penulisan kolom SQL kamu
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:update_at"`
}

func (Schedule) TableName() string {
	return "schedule"
}
