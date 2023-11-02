package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/middleware"
)

func StartServer(handlers *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	public := router.Group("api/v1")
	{
		public.POST("/register", handlers.AuthHandler.Register)
		public.POST("/login", handlers.AuthHandler.Login)
	}

	users := router.Group("api/v1/users").Use(middleware.ValidateHeader(), middleware.IsAdmin())
	{
		users.GET("/", handlers.AuthHandler.GetUsers)
	}

	return router
}
