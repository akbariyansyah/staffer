package employee

import (
	"backend/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"log"
)

type EmployeeController struct {
	employeeUsecase EmployeeUsecase
}

func newEmployeeController(db *pg.DB) *EmployeeController {
	return &EmployeeController{employeeUsecase: newEmployeeUsecase(db)}
}

func (ec *EmployeeController) getAllEmployees(ctx *gin.Context) {
	employees, err := ec.employeeUsecase.getAllEmployees()
	if err != nil {
		ctx.JSON(504, "Failed, bda request")
	}
	ctx.JSON(200, gin.H{
		"code":   200,
		"status": "ok",
		"result": employees,
	})
}
func (ec *EmployeeController) createEmployee(ctx *gin.Context) {
	var employee Employee

	err := ctx.BindJSON(&employee)
	if err != nil {
		panic(err)
	}
	err = helper.ValidateRequest(employee)
	if err != nil {
		ctx.JSON(400, "cannot be empty")
		return
	}
	err = ec.employeeUsecase.createEmployee(&employee)
	if err != nil {
		ctx.JSON(504, "Internal server error")
		return
	}
	ctx.JSON(200, gin.H{
		"code":   200,
		"status": "ok",
	})
}
func (ec *EmployeeController) updateEmployee(ctx *gin.Context) {
	var employee Employee

	err := ctx.BindJSON(&employee)

	if err != nil {
		log.Println(err)
	}
	err = helper.ValidateRequest(employee)
	if err != nil {
		ctx.JSON(400, "cannot be empty")
		return
	}
	err = ec.employeeUsecase.updateEmployee(&employee)
	if err != nil {
		log.Println(err)
		ctx.JSON(504, "internal server error")
		return
	}
	ctx.JSON(200, gin.H{
		"code":   200,
		"status": "ok",
	})
}
func (ec *EmployeeController) deleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ec.employeeUsecase.deleteEmployee(&id)

	if err != nil {
		log.Println(err)
		ctx.JSON(504, "internal server error")
		return
	}

	ctx.JSON(200, gin.H{
		"code":   200,
		"status": "ok",
	})
}
