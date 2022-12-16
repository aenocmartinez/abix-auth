package main

import (
	abixjwt "abix360/src/infraestructure/abix-jwt"
	"abix360/src/view/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func validateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "header no valid"})
		}
		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.POST("/abix360/v1/register", controller.Register)
	r.POST("/abix360/v1/login", controller.Login)

	routes := r.Group("/abix360/v1", validateHeader(), abixjwt.AuthorizeJWT())
	{
		routes.POST("/logout", controller.Logout)
		routes.PUT("/reset-password", controller.ResetPassword)
	}

	r.Run(":8080")
}
