package employee

import (
	"log"
	"net/http"
	"staffer/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Controller -> this type connecting controller and usecase layer through interface
type Controller struct {
	eu IEmployeeUsecase
}

func NewEmployeeController(e *echo.Echo, service IEmployeeUsecase) {
	controller := Controller{eu: service}
	e.GET("/employee", controller.GetEmployees)
	e.POST("/employee", controller.CreateEmployee)
	e.PUT("/employee", controller.UpdateEmployee)
	e.DELETE("/employee/:id", controller.DeleteEmployee)
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
func (ec Controller) CreateEmployee(ctx echo.Context) error {
	request := new(model.Employee)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(504, "Bad request ")
	}
	err := ec.eu.CreateEmployee(request)
	if err != nil {
		log.Println(err)
		ctx.JSON(501, "internal server error")
	}
	return nil
}
func (ec Controller) UpdateEmployee(ctx echo.Context) error {
	request := new(model.Employee)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(504, "Bad request ")
	}
	err := ec.eu.UpdateEmployee(request)
	if err != nil {
		log.Println(err)
		ctx.JSON(501, "internal server error")
	}
	return nil
}
func (ec Controller) DeleteEmployee(ctx echo.Context) error {
	id := ctx.Param("id")
	err := ec.eu.DeleteEmployee(id)
	if err != nil {
		return ctx.JSON(504, "bad request")
	}
	return ctx.JSON(200, "success")
}
