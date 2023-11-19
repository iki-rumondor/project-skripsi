package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/middleware"
)

func StartServer(handlers *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12,
	}))

	// router.Use(cors.Default())

	public := router.Group("api/v1")
	{
		public.POST("register", handlers.AuthHandler.Register)
		public.POST("login", handlers.AuthHandler.Login)
		public.GET("jurusan", handlers.UtilHandler.GetAllJurusan)
		
	}

	admin := router.Group("api/v1").Use(middleware.ValidateHeader(), middleware.IsAdmin())
	{
		admin.GET("/", handlers.AuthHandler.GetUsers)
		admin.GET("prodi", handlers.ProdiHandler.GetAllProdi)
		admin.POST("prodi", handlers.ProdiHandler.CreateProdi)
		admin.GET("prodi/:id", handlers.ProdiHandler.GetProdiByID)
		admin.DELETE("prodi/:id", handlers.ProdiHandler.DeleteProdi)
		admin.PUT("prodi/:id", handlers.ProdiHandler.UpdateProdi)
	}	

	return router
}
