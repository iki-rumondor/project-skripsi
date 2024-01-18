package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type ScheduleRepoImplementation struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepoImplementation{
		db: db,
	}
}

func (r *ScheduleRepoImplementation) CreateSchedule(model *domain.Schedule) error {
	return r.db.Create(&model).Error
}

func (r *ScheduleRepoImplementation) FindAllSchedule() (*[]domain.Schedule, error) {
	var result []domain.Schedule
	if err := r.db.Preload("AssessmentType").Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ScheduleRepoImplementation) FindScheduleByUuid(uuid string) (*domain.Schedule, error) {
	var result domain.Schedule
	if err := r.db.Preload("AssessmentType").First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ScheduleRepoImplementation) UpdateSchedule(model *domain.Schedule) error {
	return r.db.Model(&model).Updates(&model).Error
}

func (r *ScheduleRepoImplementation) DeleteSchedule(model *domain.Schedule) error {
	return r.db.Delete(&model).Error
}
