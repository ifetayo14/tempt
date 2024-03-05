package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
	"time"
)

var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(id string, username string) string {
	exp := time.Now().Local().Add(time.Hour * time.Duration(1)).Unix()
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      exp,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func AuthCheckHandler(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	strToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
