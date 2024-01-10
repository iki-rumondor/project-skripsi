package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
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

	public := router.Group("api")
	{
		public.POST("register", handlers.AuthHandler.Register)
		public.POST("login", handlers.AuthHandler.Login)
		public.GET("jurusan", handlers.UtilHandler.GetAllJurusan)

	}

	// admin := router.Group("api").Use(middleware.ValidateHeader(), middleware.IsAdmin())
	admin := router.Group("api")
	{
		admin.GET("/", handlers.AuthHandler.GetUsers)
		admin.GET("prodi", handlers.ProdiHandler.GetAllProdi)
		admin.GET("prodi/:id", handlers.ProdiHandler.GetProdiByID)
		admin.GET("instrumen/indikator", handlers.AdminHandler.GetAllIndikator)
		admin.GET("instrumen/instrumen-type", handlers.AdminHandler.GetAllInstrumenType)
		admin.GET("instrumen/indikator-type", handlers.AdminHandler.GetAllIndikatorType)
		admin.GET("instrumen/indikator/:id", handlers.AdminHandler.GetIndikator)

		admin.POST("prodi", handlers.ProdiHandler.CreateProdi)
		admin.POST("instrumen/indikator", handlers.AdminHandler.CreateIndikator)
		admin.POST("instrumen/instrumen-type", handlers.AdminHandler.CreateInstrumenType)
		admin.POST("instrumen/indikator-type", handlers.AdminHandler.CreateIndikatorType)

		admin.DELETE("prodi/:id", handlers.ProdiHandler.DeleteProdi)
		admin.DELETE("instrumen/instrumen-type/:id", handlers.AdminHandler.DeleteInstrumenType)
		admin.DELETE("instrumen/indikator-type/:id", handlers.AdminHandler.DeleteIndikatorType)

		admin.PUT("prodi/:id", handlers.ProdiHandler.UpdateProdi)
		admin.PUT("instrumen/indikator/:id", handlers.AdminHandler.UpdateIndikator)
	}

	return router
}
