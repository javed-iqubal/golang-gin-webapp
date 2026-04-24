package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PingPongMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()
		latency := time.Since(start)
		log.Printf("PingPong latency is: %s", latency)
	}
}

// auth middleware
func AuthRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		fmt.Println("Token: ", token)
		if token != "secret-token" {
			fmt.Println("Middleware authorization failed")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}
		fmt.Println("Middleware authorization pass.")
		// continue to next handler
		ctx.Next()
	}
}

func main() {

	router := gin.New()

	// global middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// route specific middleware

	router.GET("/ping", PingPongMiddleware(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong!")
	})

	// group middleware: for set of urls
	// create group
	authorized := router.Group("/")

	// attach middleware for authorized group
	authorized.Use(AuthRequiredMiddleware())
	{
		authorized.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "login")
		})
		authorized.POST("/logout", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "logout")
		})
		authorized.POST("/refresh", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "refresh")
		})
		authorized.POST("/refreah-token", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "refresh token")
		})
		authorized.POST("/register", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "register")
		})
	}

	router.Run(":8082")

}
