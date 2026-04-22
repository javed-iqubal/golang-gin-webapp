package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	reouter := gin.Default()

	reouter.GET("/welcome", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("firstname", "Guest")

		ctx.String(http.StatusOK, "Hello %s", name)
	})

	reouter.Run(":8082")
}
