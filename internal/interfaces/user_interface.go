package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type UserHandlerInterface interface {
	SignIn(*gin.Context)
}

type UserServiceInterface interface {
	VerifyUser(*request.SignIn) (string, error)
}

type UserRepoInterface interface {
	FindUserBy(column string, value interface{}) (*models.User, error)
}
