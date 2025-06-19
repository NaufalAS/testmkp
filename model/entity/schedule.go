package entity

import "mkptest/model/domain"

type JadwalEntity struct {
	JadwalID int    `json:"id"`
	MovieID   string `json:"name"`
	CinemaID  string `json:"email"`
	StartTime   string `json:"role"`
	EndTime
	ScreenNumber
	IsCancel
}

func ToJadwalEntity(jadwal domain.Schedule) JadwalEntity {
	return JadwalEntity{
		jadwalID: jadwal.ID,
		Name:   jadwal.Name,
		Email:  jadwal.Email,
		Role:   jadwal.Role,
	}
}

func ToJadwalListEntity(jadwals []domain.Schedule) []JadwalEntity {
	var jadwalData []JadwalEntity

	for _, jadwal := range jadwals {
		jadwalData = append(jadwalData, ToJadwalEntity(jadwal))

	}
	return jadwalData
}
