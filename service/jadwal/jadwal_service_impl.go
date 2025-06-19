package jadwalservice

import (
	"fmt"
	"mkptest/model/domain"
	"mkptest/model/web"
	jadwalrepo "mkptest/repository/jadwal"
	"time"
)

type JadwalServiceImpl struct {
	jadwalrepo jadwalrepo.JadwalRepo
}

func NewJadwalService(jadwalrepo jadwalrepo.JadwalRepo) *JadwalServiceImpl {
	return &JadwalServiceImpl{
		jadwalrepo: jadwalrepo,
	}
}

func (service *JadwalServiceImpl) SaveJadwal(request web.CreateScheduleRequest) (map[string]interface{}, error) {
	// Parse string waktu ke time.Time
	startTime, err := time.Parse(time.RFC3339, request.StartTime)
	if err != nil {
		return nil, fmt.Errorf("invalid start_time format: %v", err)
	}

	endTime, err := time.Parse(time.RFC3339, request.EndTime)
	if err != nil {
		return nil, fmt.Errorf("invalid end_time format: %v", err)
	}

	jadwal := domain.Schedule{
		MovieID:      request.MovieID,
		CinemaID:     request.CinemaID,
		StartTime:    startTime,
		EndTime:      endTime,
		ScreenNumber: request.ScreenNumber,
		IsCancel:     request.IsCancel,
	}

	savejadwal, errJadwal := service.jadwalrepo.SaveSchedule(jadwal)
	if errJadwal != nil {
		fmt.Println("Database error:", errJadwal)
		return nil, errJadwal
	}

	data := map[string]interface{}{
		"id":            savejadwal.ID,
		"movie_id":      savejadwal.MovieID,
		"cinema_id":     savejadwal.CinemaID,
		"start_time":    savejadwal.StartTime.Format(time.RFC3339),
		"end_time":      savejadwal.EndTime.Format(time.RFC3339),
		"screen_number": savejadwal.ScreenNumber,
		"is_cancel":     savejadwal.IsCancel,
	}
	return data, nil
}


func (service *JadwalServiceImpl) DeleteJadwal(id int) error {
	err := service.jadwalrepo.DeleteScheduleByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("jadwal dengan ID %d tidak ditemukan", id)
		}
		return err
	}
	return nil
}
