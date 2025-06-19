package jadwalservice

import "mkptest/model/web"

type JadwalService interface {
	SaveJadwal(request web.CreateScheduleRequest) (map[string]interface{}, error)
	 DeleteJadwal(id int) error
}