package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javed-iqubal/ginwebapp/04employee-management-service/handler"
)

func main() {

	router := gin.Default()

	router.GET("/employees", handler.GetEmployees)
	router.GET("/employees/:id", handler.GetEmployeeById)
	router.POST("/employees", handler.CreateEmployee)
	router.PUT("/employees/:id", handler.UpdateEmployee)
	router.DELETE("/employees/:id", handler.DeleteEmployee)
	router.Run(":8082")
}
