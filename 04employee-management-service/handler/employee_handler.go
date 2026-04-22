package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/javed-iqubal/ginwebapp/04employee-management-service/model"
)

var employees = []model.Employee{
	{ID: 1, Name: "John Doe", Age: 41, Dept: "Engineering"},
	{ID: 2, Name: "More Hugg", Age: 32, Dept: "Sale"},
	{ID: 3, Name: "Mike Salt", Age: 28, Dept: "Finance"},
}

const (
	statusError = "Error"
)

func CreateEmployee(ctx *gin.Context) {
	var employee model.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			statusError: err.Error(),
		})
		return
	}
	employee.ID = len(employees) + 1
	employees = append(employees, employee)
	ctx.JSON(http.StatusCreated, employee)

}
func GetEmployees(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, employees)
}

func GetEmployeeById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			statusError: err.Error(),
		})
		return
	}
	for _, employee := range employees {
		if employee.ID == id {
			ctx.JSON(http.StatusOK, employee)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Employee not found.",
	})
}

func UpdateEmployee(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			statusError: err.Error(),
		})
		return
	}
	var employee model.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			statusError: err.Error(),
		})
		return
	}
	for index, oldEmp := range employees {
		if oldEmp.ID == id {
			employee.ID = id
			employees[index] = employee
			ctx.JSON(http.StatusOK, employee)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Employee does not exist.",
	})
}

func DeleteEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			statusError: err.Error(),
		})
	}
	for index, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:index], employees[index+1:]...)
			ctx.Status(http.StatusNoContent)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Employee does not exist.",
	})
}
