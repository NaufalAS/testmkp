package jadwalrepo

import (
	"mkptest/model/domain"

	"gorm.io/gorm"
)

type JadwalRepositoryImpl struct {
	DB *gorm.DB
}

func NewJadwalRepository(db *gorm.DB) *JadwalRepositoryImpl {
	return &JadwalRepositoryImpl{
		DB: db,
	}
}

func (repo *JadwalRepositoryImpl) SaveSchedule(jadawal domain.Schedule) (domain.Schedule, error) {
	err := repo.DB.Create(&jadawal).Error
	if err != nil {
		return domain.Schedule{}, err
	}
	return jadawal, nil
}

func (repo *JadwalRepositoryImpl) DeleteScheduleByID(id int) error {
	result := repo.DB.Delete(&domain.Schedule{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *JadwalRepositoryImpl) GetListSchedule() ([]domain.Schedule, error) {
	var jadwal []domain.Schedule
	err := repo.DB.Find(&jadwal).Error
	if err != nil {
		return nil, err
	}
	return jadwal, nil
}

func(repo *JadwalRepositoryImpl) GetJadwalById(id int)(domain.Schedule, error){
	var jadwal domain.Schedule

	if err := repo.DB.Find(&jadwal,"id = ?", id).Error; err != nil{
		return jadwal, err
	}

	return jadwal, nil
}
