package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ScheduleRepository interface {
	CreateSchedule(model *domain.Schedule) error
	FindAllSchedule() (*[]domain.Schedule, error)
	FindScheduleByUuid(uuid string) (*domain.Schedule, error)
	UpdateSchedule(model *domain.Schedule) error
	DeleteSchedule(model *domain.Schedule) error
}
