package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Global error handler
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// process the request first
		ctx.Next()

		// check if any errors ware added to context
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}

func main() {

	router := gin.Default()

	// attache error handler global middleware
	router.Use(ErrorHandler())

	router.GET("/ok", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Okay")
	})
	router.GET("/error", func(ctx *gin.Context) {
		ctx.Error(errors.New("somthing went wrong"))
	})

	router.Run((":8082"))
}
