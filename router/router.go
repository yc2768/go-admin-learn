package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/middleware"
	"go-admin/service"
)

func App() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	r.POST("/api/v1/login", service.Login)

	return r
}
