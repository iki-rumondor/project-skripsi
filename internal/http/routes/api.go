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

	}

	department := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("DEPARTMENT"))
	{
		department.POST("subjects", handlers.SubjectHandler.CreateSubject)
		department.GET("subjects", handlers.SubjectHandler.GetAllSubjects)
		department.GET("subjects/:uuid", handlers.SubjectHandler.GetSubject)
		department.PUT("subjects/:uuid", handlers.SubjectHandler.UpdateSubject)
		department.DELETE("subjects/:uuid", handlers.SubjectHandler.DeleteSubject)
	}

	admin := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("ADMIN"))
	{
		admin.POST("departments", handlers.DepartmentHandler.CreateDepartment)
		admin.GET("departments", handlers.DepartmentHandler.GetAllDepartments)
		admin.GET("departments/:uuid", handlers.DepartmentHandler.GetDepartment)
		admin.PUT("departments/:uuid", handlers.DepartmentHandler.UpdateDepartment)
		admin.DELETE("departments/:uuid", handlers.DepartmentHandler.DeleteDepartment)
	}

	return router
}
