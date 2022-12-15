package main

import (
	"abix360/src/view/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/register", controller.Register)

	r.Run()
}
