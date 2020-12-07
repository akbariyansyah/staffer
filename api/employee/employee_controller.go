package employee

import (
	"github.com/labstack/echo/v4"
)

type EmployeeController struct {
	eu IEmployeeUsecase
}

func NewEmployeeController(e *echo.Echo, service IEmployeeUsecase) {
	controller := EmployeeController{eu: service}
	e.GET("/employee", controller.handleGetEmployees)
}
func (ec EmployeeController) handleGetEmployees(ctx echo.Context) error {
	employees,err := ec.eu.GetAllEmployees()
	if err != nil {
		return ctx.JSON(501, "internal server error")
	}
	return ctx.JSON(200, employees)
}