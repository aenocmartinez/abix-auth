package main

import (
	"abix360/src/dao"
)

func main() {

	for i := 0; i < 30; i++ {
		dao.Instance()
	}
	// r := gin.Default()
	// r.POST("/register", controller.Register)

	// r.Run()
}
