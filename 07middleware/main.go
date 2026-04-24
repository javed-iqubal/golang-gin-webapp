package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// router := gin.Default()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong!")
	})

	router.Run(":8082")
}
