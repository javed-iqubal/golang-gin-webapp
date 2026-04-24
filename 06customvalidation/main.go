package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// define struct
type UserRequest struct {
	Username string `json:"username" binding:"required,is-admin"`
	Age      int    `json:"age" binding:"required,gt=18"`
}

// define custom validation logic
var validateAdmin validator.Func = func(fl validator.FieldLevel) bool {
	// get field value
	value := fl.Field().String()
	fmt.Println(value)
	return strings.Contains(strings.ToLower(value), "admin")
}

func main() {

	router := gin.Default()

	// register custom validator with Gin engine
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is-admin", validateAdmin)

	}

	router.POST("/register", func(ctx *gin.Context) {
		var user UserRequest
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "validation failed",
				"details": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to " + user.Username,
		})
	})

	router.Run(":8082")

}
