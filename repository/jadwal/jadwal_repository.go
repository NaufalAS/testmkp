package jadwalrepo

import "mkptest/model/domain"

type JadwalRepo interface {
	SaveSchedule(jadawal domain.Schedule) (domain.Schedule, error)
	DeleteScheduleByID(id int) error
	GetListSchedule() ([]domain.Schedule, error)
	GetJadwalById(id int)(domain.Schedule, error)
}