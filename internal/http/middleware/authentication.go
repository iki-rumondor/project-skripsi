package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/utils"
)

func IsValidJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerToken = c.Request.Header.Get("Authorization")
		var bearer = strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			utils.HandleError(c, response.UNAUTH_ERR("Bearer Token Tidak Ditemukan"))
			return
		}

		jwt := strings.Split(headerToken, " ")[1]

		mapClaims, err := utils.VerifyToken(jwt)
		if err != nil {
			utils.HandleError(c, response.UNAUTH_ERR("Token Tidak Valid"))
			return
		}

		c.Set("map_claims", mapClaims)
		c.Next()

	}
}

func SetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		id, ok := mapClaims["user_id"].(float64)
		if !ok {
			utils.HandleError(c, response.UNAUTH_ERR("Token Tidak Valid"))
			return
		}

		c.Set("user_id", uint(id))
		c.Next()

	}
}

func SetUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		role, ok := mapClaims["role"].(string)
		if !ok {
			utils.HandleError(c, response.UNAUTH_ERR("Token Tidak Valid"))
			return
		}

		c.Set("user_role", role)
		c.Next()

	}
}
