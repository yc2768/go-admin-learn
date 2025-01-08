package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE")
		c.Header("Access-Control-Allow-Headers", "origin, X-Requested-With, Content-Type, Accept, Authorization,AccessToken,userToken")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control,Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
