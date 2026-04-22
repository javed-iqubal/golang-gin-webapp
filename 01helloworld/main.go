package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// initialize the Gin Engine for middleware support
	router := gin.Default()

	// router with home page
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Go learner, Welcome to learn Gin Framework",
		})
	})
	fmt.Println("Gin Server is running...")
	// default port is 8080
	// router.Run()

	router.Run(":8082")

}
