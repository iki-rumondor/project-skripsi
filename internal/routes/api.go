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
		admin.GET("instrumen/indikator", handlers.InstrumenHandler.GetAllIndikator)
		admin.GET("instrumen/instrumen-type", handlers.InstrumenHandler.GetAllInstrumenType)
		admin.GET("instrumen/indikator-type", handlers.InstrumenHandler.GetAllIndikatorType)
		admin.GET("instrumen/indikator/:id", handlers.InstrumenHandler.GetIndikator)

		admin.POST("prodi", handlers.ProdiHandler.CreateProdi)
		admin.POST("instrumen/indikator", handlers.InstrumenHandler.CreateIndikator)
		admin.POST("instrumen/instrumen-type", handlers.InstrumenHandler.CreateInstrumenType)
		admin.POST("instrumen/indikator-type", handlers.InstrumenHandler.CreateIndikatorType)

		admin.DELETE("prodi/:id", handlers.ProdiHandler.DeleteProdi)
		admin.DELETE("instrumen/instrumen-type/:id", handlers.InstrumenHandler.DeleteInstrumenType)
		admin.DELETE("instrumen/indikator-type/:id", handlers.InstrumenHandler.DeleteIndikatorType)

		admin.PUT("prodi/:id", handlers.ProdiHandler.UpdateProdi)
		admin.PUT("instrumen/indikator/:id", handlers.InstrumenHandler.UpdateIndikator)

		admin.POST("assessments/type", handlers.AssTypeHandler.CreateAssType)
		admin.GET("assessments/type", handlers.AssTypeHandler.GetAllAssType)
		admin.GET("assessments/type/:uuid", handlers.AssTypeHandler.GetAssTypeByUuid)
		admin.PUT("assessments/type/:uuid", handlers.AssTypeHandler.UpdateAssType)
		admin.DELETE("assessments/type/:uuid", handlers.AssTypeHandler.DeleteAssType)

		admin.POST("assessments/question", handlers.AssQuestionHandler.CreateAssQuestion)
		admin.GET("assessments/question", handlers.AssQuestionHandler.GetAllAssQuestion)
		admin.GET("assessments/question/:uuid", handlers.AssQuestionHandler.GetAssQuestionByUuid)
		admin.PUT("assessments/question/:uuid", handlers.AssQuestionHandler.UpdateAssQuestion)
		admin.DELETE("assessments/question/:uuid", handlers.AssQuestionHandler.DeleteAssQuestion)

		admin.POST("response", handlers.ResponseHandler.CreateResponse)
		admin.GET("response", handlers.ResponseHandler.GetAllResponse)
		admin.GET("response/:uuid", handlers.ResponseHandler.GetResponseByUuid)
		admin.PUT("response/:uuid", handlers.ResponseHandler.UpdateResponse)
		admin.DELETE("response/:uuid", handlers.ResponseHandler.DeleteResponse)
	}

	return router
}
