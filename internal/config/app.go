package config

import (
	"github.com/iki-rumondor/go-monev/internal/http/handlers"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/repository"
	"github.com/iki-rumondor/go-monev/internal/services"
	"gorm.io/gorm"
)

type Handlers struct {
	UserHandler              interfaces.UserHandlerInterface
	DepartmentHandler        interfaces.DepartmentHandlerInterface
	SubjectHandler           interfaces.SubjectHandlerInterface
	MajorHandler             interfaces.MajorHandlerInterface
	LaboratoryHandler        interfaces.LaboratoryHandlerInterface
	AcademicYearHandler      interfaces.AcademicYearHandlerInterface
	AcademicPlanHandler      interfaces.AcademicPlanHandlerInterface
	PracticalModuleHandler   interfaces.PracticalModuleHandlerInterface
	PracticalToolHandler     interfaces.PracticalToolHandlerInterface
	TeacherHandler           interfaces.TeacherHandlerInterface
	TeacherSkillHandler      interfaces.TeacherSkillHandlerInterface
	FacilityHandler          interfaces.FacilityHandlerInterface
	FacilityConditionHandler interfaces.FacilityConditionHandlerInterface
	MiddleMonevHandler       interfaces.MiddleMonevHandlerInterface
	LastMonevHandler         interfaces.LastMonevHandlerInterface
	PdfHandler               interfaces.PdfHandlerInterface
}

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

	laboratory_repo := repository.NewLaboratoryRepository(db)
	laboratory_service := services.NewLaboratoryService(laboratory_repo)
	laboratory_handler := handlers.NewLaboratoryHandler(laboratory_service)

	teacher_repo := repository.NewTeacherRepository(db)
	teacher_service := services.NewTeacherService(teacher_repo)
	teacher_handler := handlers.NewTeacherHandler(teacher_service)

	academic_year_repo := repository.NewAcademicYearRepository(db)
	academic_year_service := services.NewAcademicYearService(academic_year_repo)
	academic_year_handler := handlers.NewAcademicYearHandler(academic_year_service)

	academic_plan_repo := repository.NewAcademicPlanRepository(db)
	academic_plan_service := services.NewAcademicPlanService(academic_plan_repo)
	academic_plan_handler := handlers.NewAcademicPlanHandler(academic_plan_service)

	practical_module_repo := repository.NewPracticalModuleRepository(db)
	practical_module_service := services.NewPracticalModuleService(practical_module_repo)
	practical_module_handler := handlers.NewPracticalModuleHandler(practical_module_service)

	practical_tool_repo := repository.NewPracticalToolRepository(db)
	practical_tool_service := services.NewPracticalToolService(practical_tool_repo)
	practical_tool_handler := handlers.NewPracticalToolHandler(practical_tool_service)

	teacher_skill_repo := repository.NewTeacherSkillRepository(db)
	teacher_skill_service := services.NewTeacherSkillService(teacher_skill_repo)
	teacher_skill_handler := handlers.NewTeacherSkillHandler(teacher_skill_service)

	facility_repo := repository.NewFacilityRepository(db)
	facility_service := services.NewFacilityService(facility_repo)
	facility_handler := handlers.NewFacilityHandler(facility_service)

	facility_condition_repo := repository.NewFacilityConditionRepository(db)
	facility_condition_service := services.NewFacilityConditionService(facility_condition_repo)
	facility_condition_handler := handlers.NewFacilityConditionHandler(facility_condition_service)

	middle_monev_repo := repository.NewMiddleMonevRepository(db)
	middle_monev_service := services.NewMiddleMonevService(middle_monev_repo)
	middle_monev_handler := handlers.NewMiddleMonevHandler(middle_monev_service)

	last_monev_repo := repository.NewLastMonevRepository(db)
	last_monev_service := services.NewLastMonevService(last_monev_repo)
	last_monev_handler := handlers.NewLastMonevHandler(last_monev_service)

	pdf_repo := repository.NewPdfRepository(db)
	pdf_service := services.NewPdfService(pdf_repo)
	pdf_handler := handlers.NewPdfHandler(pdf_service)

	return &Handlers{
		UserHandler:              user_handler,
		DepartmentHandler:        department_handler,
		SubjectHandler:           subject_handler,
		MajorHandler:             major_handler,
		LaboratoryHandler:        laboratory_handler,
		AcademicYearHandler:      academic_year_handler,
		AcademicPlanHandler:      academic_plan_handler,
		PracticalModuleHandler:   practical_module_handler,
		PracticalToolHandler:     practical_tool_handler,
		TeacherHandler:           teacher_handler,
		TeacherSkillHandler:      teacher_skill_handler,
		FacilityHandler:          facility_handler,
		FacilityConditionHandler: facility_condition_handler,
		MiddleMonevHandler:       middle_monev_handler,
		LastMonevHandler:         last_monev_handler,
		PdfHandler:               pdf_handler,
	}
}
