package middleware

import (
	"agit/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.AuthCheckHandler(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.JSONResponse{
				Message: "Unauthorized",
				Errors:  err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
