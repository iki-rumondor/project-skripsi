package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/config"
	"github.com/iki-rumondor/go-monev/internal/http/middleware"
)

func StartServer(handlers *config.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12,
	}))

	public := router.Group("api")
	{
		public.POST("signin", handlers.UserHandler.SignIn)
		public.GET("chart/departments/year/:yearUuid", handlers.UserHandler.GetDepartmentsChart)
		public.GET("academic-years", handlers.AcademicYearHandler.GetAllAcademicYears)
	}
	
	user := router.Group("api").Use(middleware.IsValidJWT())
	{
		user.GET("academic-years/:uuid", handlers.AcademicYearHandler.GetAcademicYear)
		user.GET("dashboards", handlers.UserHandler.GetDashboardAdmin)
		user.GET("settings", handlers.UserHandler.GetSettings)

		user.GET("departments", handlers.DepartmentHandler.GetAllDepartments)
		user.GET("monev/departments/:departmentUuid/years/:yearUuid", handlers.UserHandler.GetDepartmentMonev)
		user.GET("academic-plans/departments/:departmentUuid/years/:yearUuid", handlers.AcademicPlanHandler.GetDepartment)
		user.GET("practical-modules/departments/:departmentUuid/years/:yearUuid", handlers.PracticalModuleHandler.GetByDepartment)
		user.GET("practical-tools/departments/:departmentUuid/years/:yearUuid", handlers.PracticalToolHandler.GetByDepartment)
		user.GET("teacher-skills/departments/:departmentUuid/years/:yearUuid", handlers.TeacherSkillHandler.GetByDepartment)
		user.GET("facility-conditions/departments/:departmentUuid/years/:yearUuid", handlers.FacilityConditionHandler.GetByDepartment)
		user.GET("teacher-attendences/departments/:departmentUuid/years/:yearUuid", handlers.MiddleMonevHandler.GetTeacherAttendencesByDepartment)
		user.GET("student-attendences/departments/:departmentUuid/years/:yearUuid", handlers.MiddleMonevHandler.GetStudentAttendencesByDepartment)

		user.GET("report/:typeReport/departments/:departmentUuid/years/:yearUuid", handlers.PdfHandler.CreateReport)
	}

	department := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("DEPARTMENT"), middleware.SetUserUuid())
	{
		department.POST("subjects", handlers.SubjectHandler.CreateSubject)
		department.GET("subjects", handlers.SubjectHandler.GetAllSubjects)
		department.GET("subjects/:uuid", handlers.SubjectHandler.GetSubject)
		department.PUT("subjects/:uuid", handlers.SubjectHandler.UpdateSubject)
		department.DELETE("subjects/:uuid", handlers.SubjectHandler.DeleteSubject)

		department.POST("laboratories", handlers.LaboratoryHandler.CreateLaboratory)
		department.GET("laboratories", handlers.LaboratoryHandler.GetAllLaboratories)
		department.GET("laboratories/:uuid", handlers.LaboratoryHandler.GetLaboratory)
		department.PUT("laboratories/:uuid", handlers.LaboratoryHandler.UpdateLaboratory)
		department.DELETE("laboratories/:uuid", handlers.LaboratoryHandler.DeleteLaboratory)

		department.POST("facilities", handlers.FacilityHandler.CreateFacility)
		department.GET("facilities", handlers.FacilityHandler.GetAllFacilities)
		department.GET("facilities/:uuid", handlers.FacilityHandler.GetFacility)
		department.PUT("facilities/:uuid", handlers.FacilityHandler.UpdateFacility)
		department.DELETE("facilities/:uuid", handlers.FacilityHandler.DeleteFacility)

		department.POST("teachers", handlers.TeacherHandler.CreateTeacher)
		department.GET("teachers", handlers.TeacherHandler.GetAllTeachers)
		department.GET("teachers/:uuid", handlers.TeacherHandler.GetTeacher)
		department.PUT("teachers/:uuid", handlers.TeacherHandler.UpdateTeacher)
		department.DELETE("teachers/:uuid", handlers.TeacherHandler.DeleteTeacher)

		department.POST("academic-plans", handlers.AcademicPlanHandler.CreateAcademicPlan)
		department.GET("academic-plans/years/:yearUuid", handlers.AcademicPlanHandler.GetAllAcademicPlans)
		department.GET("academic-plans/:uuid", handlers.AcademicPlanHandler.GetAcademicPlan)
		department.PUT("academic-plans/:uuid", handlers.AcademicPlanHandler.UpdateAcademicPlan)
		department.DELETE("academic-plans/:uuid", handlers.AcademicPlanHandler.DeleteAcademicPlan)

		department.GET("academic-plans/middle/years/:yearUuid", handlers.AcademicPlanHandler.GetMiddle)
		department.PATCH("academic-plans/:uuid/middle", handlers.AcademicPlanHandler.UpdateMiddle)
		department.PATCH("academic-plans/:uuid/last", handlers.AcademicPlanHandler.UpdateLast)

		department.POST("practical-modules", handlers.PracticalModuleHandler.CreatePracticalModule)
		department.GET("practical-modules/years/:yearUuid", handlers.PracticalModuleHandler.GetAllPracticalModules)
		department.GET("practical-modules/:uuid", handlers.PracticalModuleHandler.GetPracticalModule)
		department.PUT("practical-modules/:uuid", handlers.PracticalModuleHandler.UpdatePracticalModule)
		department.DELETE("practical-modules/:uuid", handlers.PracticalModuleHandler.DeletePracticalModule)

		department.POST("practical-tools", handlers.PracticalToolHandler.CreatePracticalTool)
		department.GET("practical-tools/years/:yearUuid", handlers.PracticalToolHandler.GetAllPracticalTools)
		department.GET("practical-tools/:uuid", handlers.PracticalToolHandler.GetPracticalTool)
		department.PUT("practical-tools/:uuid", handlers.PracticalToolHandler.UpdatePracticalTool)
		department.DELETE("practical-tools/:uuid", handlers.PracticalToolHandler.DeletePracticalTool)

		department.POST("teacher-skills", handlers.TeacherSkillHandler.CreateTeacherSkill)
		department.GET("teacher-skills", handlers.TeacherSkillHandler.GetAllTeacherSkills)
		department.GET("teacher-skills/years/:yearUuid", handlers.TeacherSkillHandler.GetByYear)
		department.GET("teacher-skills/:uuid", handlers.TeacherSkillHandler.GetTeacherSkill)
		department.PUT("teacher-skills/:uuid", handlers.TeacherSkillHandler.UpdateTeacherSkill)
		department.DELETE("teacher-skills/:uuid", handlers.TeacherSkillHandler.DeleteTeacherSkill)

		department.POST("facility-conditions", handlers.FacilityConditionHandler.CreateFacilityCondition)
		department.GET("facility-conditions/years/:yearUuid", handlers.FacilityConditionHandler.GetFacilityConditionsByYear)
		department.GET("facility-conditions/years/:yearUuid/options", handlers.FacilityConditionHandler.GetFacilityOptions)
		department.GET("facility-conditions/:uuid", handlers.FacilityConditionHandler.GetFacilityCondition)
		department.PUT("facility-conditions/:uuid", handlers.FacilityConditionHandler.UpdateFacilityCondition)
		department.DELETE("facility-conditions/:uuid", handlers.FacilityConditionHandler.DeleteFacilityCondition)

		department.GET("middle-monev/years/:yearUuid", handlers.MiddleMonevHandler.CountTotalMonev)
		department.POST("middle-monev/teacher-attendences", handlers.MiddleMonevHandler.CreateTeacherAttendence)
		department.GET("middle-monev/teacher-attendences/years/:yearUuid", handlers.MiddleMonevHandler.GetTeacherAttendences)
		department.PATCH("middle-monev/last/teacher-attendences/:uuid", handlers.MiddleMonevHandler.UpdateLastTeacherAttendence)
		department.PATCH("middle-monev/last/student-attendences/:uuid", handlers.MiddleMonevHandler.UpdateLastStudentAttendence)
		department.DELETE("middle-monev/teacher-attendences/:uuid", handlers.MiddleMonevHandler.DeleteTeacherAttendence)

		department.POST("middle-monev/student-attendences", handlers.MiddleMonevHandler.CreateStudentAttendence)
		department.GET("middle-monev/student-attendences/years/:yearUuid", handlers.MiddleMonevHandler.GetStudentAttendences)
		department.DELETE("middle-monev/student-attendences/:uuid", handlers.MiddleMonevHandler.DeleteStudentAttendence)

		department.GET("last-monev/years/:yearUuid", handlers.LastMonevHandler.CountLastMonev)
		department.PATCH("last-monev/student-passed/:uuid", handlers.LastMonevHandler.UpdateStudentPass)
		department.PATCH("last-monev/student-final/:uuid", handlers.LastMonevHandler.UpdateStudentFinal)
		department.PATCH("last-monev/grade/teacher-attendences/:uuid", handlers.LastMonevHandler.UpdateTeacherGrade)

		department.GET("subjects/teacher-attendences/years/:yearUuid", handlers.SubjectHandler.GetTeacherAttendenceSubjects)
		department.GET("subjects/student-attendences/years/:yearUuid", handlers.SubjectHandler.GetStudentAttendenceSubjects)
		department.GET("subjects/tables/:tableName/years/:yearUuid", handlers.SubjectHandler.GetOuterSubjects)

		department.GET("subjects/practical", handlers.SubjectHandler.GetAllPracticalSubjects)
		department.GET("users/first-monev/years/:yearUuid", handlers.UserHandler.CountMonevByYear)
		department.GET("users/:userUuid", handlers.UserHandler.GetDepartmentData)
	}

	admin := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("ADMIN"))
	{
		admin.POST("majors", handlers.MajorHandler.CreateMajor)
		admin.GET("majors", handlers.MajorHandler.GetAllMajors)
		admin.GET("majors/:uuid", handlers.MajorHandler.GetMajor)
		admin.PUT("majors/:uuid", handlers.MajorHandler.UpdateMajor)
		admin.DELETE("majors/:uuid", handlers.MajorHandler.DeleteMajor)

		admin.POST("departments", handlers.DepartmentHandler.CreateDepartment)
		admin.GET("departments/:uuid", handlers.DepartmentHandler.GetDepartment)
		admin.PUT("departments/:uuid", handlers.DepartmentHandler.UpdateDepartment)
		admin.DELETE("departments/:uuid", handlers.DepartmentHandler.DeleteDepartment)

		admin.POST("academic-years", handlers.AcademicYearHandler.CreateAcademicYear)
		admin.PUT("academic-years/:uuid", handlers.AcademicYearHandler.UpdateAcademicYear)
		admin.PATCH("academic-years/:uuid/open", handlers.AcademicYearHandler.UpdateOpen)
		admin.DELETE("academic-years/:uuid", handlers.AcademicYearHandler.DeleteAcademicYear)

		admin.PUT("academic-years/:uuid/monev", handlers.AcademicYearHandler.UpdateTimeMonev)

		admin.PATCH("settings/step", handlers.UserHandler.UpdateStepMonev)

		admin.POST("users", handlers.UserHandler.CreateUser)
		admin.GET("users", handlers.UserHandler.GetUsers)
		admin.GET("users/roles", handlers.UserHandler.GetRoles)

	}

	return router
}
