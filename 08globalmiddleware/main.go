package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// router specific middleware - pingpongmiddleware
func PingPongMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Next() // very important
		latency := time.Since(t)
		log.Printf("PingPong latency: %s", latency)
	}
}

func main() {

	router := gin.New()

	// Global middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// router specific middleware
	// router.method("url","middleware",handler)
	router.GET("/ping", PingPongMiddleware(), func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, "Pong!")

	})

	router.Run(":8082")
}
