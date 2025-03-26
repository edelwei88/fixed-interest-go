package main

import (
	initial "github.com/edelwei88/fixed-interest-go/initial"
	"github.com/gin-gonic/gin"
)

func initialize() {
	initial.LoadEnv()
	initial.ConnectToDB()
}

func main() {
	initialize()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
