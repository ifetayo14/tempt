package router

import (
	"agit/handler"
	"agit/middleware"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {

	r := gin.Default()

	// public access endpoint
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	// required jwt endpoint
	guardedJWT := r.Group("").Use(middleware.Authentication())

	guardedJWT.POST("/employee", handler.Create)
	guardedJWT.PATCH("/employee", handler.Update)
	guardedJWT.GET("/employee", handler.GetAll)
	guardedJWT.GET("/employee/:id", handler.Detail)
	guardedJWT.DELETE("/employee/:id", handler.Delete)

	return r
}
