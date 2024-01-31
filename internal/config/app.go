package config

import (
	"github.com/iki-rumondor/go-monev/internal/http/handlers"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/repository"
	"github.com/iki-rumondor/go-monev/internal/services"
	"gorm.io/gorm"
)

type Handlers struct {
	UserHandler       interfaces.UserHandlerInterface
	DepartmentHandler interfaces.DepartmentHandlerInterface
	SubjectHandler    interfaces.SubjectHandlerInterface
	MajorHandler      interfaces.MajorHandlerInterface
}

// type Handlers struct {
// 	UserHandler       *handlers.UserHandler
// 	DepartmentHandler *handlers.DepartmentHandler
// 	SubjectHandler    *handlers.SubjectHandler
// }

func GetAppHandlers(db *gorm.DB) *Handlers {

	user_repo := repository.NewUserRepository(db)
	user_service := services.NewUserService(user_repo)
	user_handler := handlers.NewUserHandler(user_service)

	major_repo := repository.NewMajorRepository(db)
	major_service := services.NewMajorService(major_repo)
	major_handler := handlers.NewMajorHandler(major_service)

	department_repo := repository.NewDepartmentRepository(db)
	department_service := services.NewDepartmentService(department_repo)
	department_handler := handlers.NewDepartmentHandler(department_service)

	subject_repo := repository.NewSubjectRepository(db)
	subject_service := services.NewSubjectService(subject_repo)
	subject_handler := handlers.NewSubjectHandler(subject_service)

	return &Handlers{
		UserHandler:       user_handler,
		DepartmentHandler: department_handler,
		SubjectHandler:    subject_handler,
		MajorHandler:      major_handler,
	}
}
