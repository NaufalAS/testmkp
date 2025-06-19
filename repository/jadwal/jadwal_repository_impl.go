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
