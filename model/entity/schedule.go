package entity

import (
	"mkptest/model/domain"
	"time"
)

type JadwalEntity struct {
	JadwalID int    `json:"id"`
	MovieID   int `json:"name"`
	CinemaID  int `json:"email"`
	StartTime   time.Time `json:"role"`
	ScreenNumber int `json:"screen_number"`
	IsCancel  bool `json:"is_cancle"`
}

func ToJadwalEntity(jadwal domain.Schedule) JadwalEntity {
	return JadwalEntity{
		JadwalID: jadwal.ID,
		MovieID:   jadwal.MovieID,
		CinemaID:  jadwal.CinemaID,
		StartTime:   jadwal.StartTime,
		ScreenNumber: jadwal.ScreenNumber,
		IsCancel: jadwal.IsCancel,
	}
}

func ToJadwalListEntity(jadwals []domain.Schedule) []JadwalEntity {
	var jadwalData []JadwalEntity

	for _, jadwal := range jadwals {
		jadwalData = append(jadwalData, ToJadwalEntity(jadwal))

	}
	return jadwalData
}
