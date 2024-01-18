package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

var secretKey = "fabsence"

type JwtClaims struct {
	UserUuid string   `json:"uuid"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(UserUuid string, role string) (string, error) {

	expireTime := time.Now().Add(time.Duration(1) * 24 * time.Hour)
	claims := &JwtClaims{
		UserUuid: UserUuid,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(strToken string) (jwt.MapClaims, error) {
	errResponse := &response.Error{
		Code:    403,
		Message: "Token anda tidak valid",
	}

	token, _ := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	mapClaims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errResponse
	}

	return mapClaims, nil
}
