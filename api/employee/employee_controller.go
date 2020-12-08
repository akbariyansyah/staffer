package employee

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

// Controller -> this type connecting controller and usecase layer through interface
type Controller struct {
	eu IEmployeeUsecase
}

func NewEmployeeController(e *echo.Echo, service IEmployeeUsecase) {
	controller := Controller{eu: service}
	e.GET("/employee", controller.GetEmployees)
	e.POST("/employee", controller.CreateEmployee)
}
func (ec Controller) GetEmployees(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")
	if page == "" || limit == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "error bad request",
		})
	}
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	employees, err := ec.eu.GetAllEmployees(pageInt, limitInt)

	if err != nil {
		log.Println(err)
		return ctx.JSON(501, "internal server error")
	}
	return ctx.JSON(200, employees)
}
func (ec Controller) CreateEmployee(cxt echo.Context) error {
	return nil
}
