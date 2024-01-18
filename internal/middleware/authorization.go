package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetString("jwt")
		if jwt == "" {
			utils.HandleError(c, &response.Error{
				Code:    403,
				Message: "Token Anda Tidak Valid",
			})
			return
		}

		mapClaims, err := utils.VerifyToken(jwt)
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		role := mapClaims["role"].(string)
		if role != "ADMIN" {
			utils.HandleError(c, &response.Error{
				Code:    403,
				Message: "Token Anda Tidak Valid",
			})
			return
		}
		c.Next()

	}
}
