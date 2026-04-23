package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=5,max=10,alphanum"`
}

func LoginHandler(ctx *gin.Context) {

	var body Login
	/**
	MustBindJson -
	1. Parse the json body into the body struct
	- if it fails(e.g- missing field or invalid json), it set error message
	*/

	if err := ctx.BindJSON(&body); err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "success",
		"data":   body.Username,
	})

}

func LoginHandlerWithError(ctx *gin.Context) {
	var body Login
	if err := ctx.ShouldBind(&body); err != nil {
		var validator validator.ValidationErrors
		if errors.As(err, &validator) {
			out := make(map[string]string)
			for _, ferr := range validator {
				out[ferr.Field()] = getFriendlyError(ferr)
			}
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"errors": out,
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json format.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "logged In!",
	})

}

func getFriendlyError(fieldError validator.FieldError) string {

	switch fieldError.Tag() {
	case "required":
		return "This field is required."
	case "max":
		return fmt.Sprintf("The maximum length field must be less than %s", fieldError.Param())
	case "min":
		return fmt.Sprintf("The minimum length field must be greater than %s", fieldError.Param())
	case "email":
		return "The email address is required."
	case "alphanum":
		return "The field must be alpha numeric"
	}
	return "Invalid value"
}

func main() {

	router := gin.Default()
	router.POST("/login", LoginHandler)
	router.POST("/logincustom", LoginHandlerWithError)
	router.Run(":8082")
}
