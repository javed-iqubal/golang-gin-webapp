package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":  gin.H{"email": "foo@bar.com", "phone": 1234567890},
	"subu": gin.H{"email": "subu@bar.com", "phone": 1234000000},
	"lena": gin.H{"email": "lena@bar.com", "phone": 1234999999},
}

func main() {

	router := gin.Default()

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"subu": "1234",
		"lena": "lena",
	}))

	// end points - /admin/secrets
	authorized.GET("/secrets", func(ctx *gin.Context) {
		// get user, it was set by th BasicAuth middleware

		user := ctx.MustGet(gin.AuthUserKey).(string)

		if secret, ok := secrets[user]; ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "No secret"})
		}
	})

	router.Run((":8082"))

}
