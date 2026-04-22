package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/users", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": " Users GET",
		})
	})

	router.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"message": "User GET by param " + id,
		})
	})

	router.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "User POST by param ",
		})
	})

	router.PUT("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"message": "User PUT by param " + id,
		})
	})

	router.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"message": "User DELETE by param " + id,
		})
	})

	fmt.Println("Gin Server is running...")
	router.Run(":8082")

}
