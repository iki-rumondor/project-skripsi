package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

func ValidateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerToken = c.Request.Header.Get("Authorization")
		var bearer = strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			utils.HandleError(c, &response.Error{
				Code:    403,
				Message: "Bearer Token Tidak Valid",
			})
			return
		}

		stringToken := strings.Split(headerToken, " ")[1]

		c.Set("jwt", stringToken)
		c.Next()
	}
}
