package employee

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// EmployeeController -> this type connecting controller and usecase layer through interface
type EmployeeController struct {
	eu IEmployeeUsecase
}

func NewEmployeeController(e *echo.Echo, service IEmployeeUsecase) {
	controller := EmployeeController{eu: service}
	e.GET("/employee", controller.HandleGetEmployees)
	e.POST("/employee", controller.HandleCreateEmployee)
}
func (ec EmployeeController) HandleGetEmployees(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")
	if page == "" || limit == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "error bad request",
		})
	}
	employees, err := ec.eu.GetAllEmployees(page, limit)
	
	if err != nil {
		log.Println(err)
		return ctx.JSON(501, "internal server error")
	}
	return ctx.JSON(200, employees)
}
func (ec EmployeeController) HandleCreateEmployee(cxt echo.Context) error {
	return nil
}
