package jadwalservice

import (
	"mkptest/model/entity"
	"mkptest/model/web"
)

type JadwalService interface {
	SaveJadwal(request web.CreateScheduleRequest) (map[string]interface{}, error)
	 DeleteJadwal(id int) error
	 GetJadwalList() ([]entity.JadwalEntity, error) 
	 GetJadwalById(id int) (entity.JadwalEntity, error)
}