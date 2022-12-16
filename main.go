package main

import (
	"abix360/src/view/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)

	r.Run(":8080")
}
