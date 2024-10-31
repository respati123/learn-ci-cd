package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "masukk")
	})

	app.Run(":8000")
}
