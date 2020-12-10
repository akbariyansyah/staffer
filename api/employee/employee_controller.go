package employee

import (
	"log"
	"net/http"
	"staffer/model"
	"strconv"

	"github.com/labstack/echo"
)

// Controller this type connecting controller and usecase layer through interface
type Controller struct {
	eu IEmployeeUsecase
}

func NewEmployeeController(e *echo.Echo, service IEmployeeUsecase) {
	controller := Controller{eu: service}
	e.GET("/employee", controller.GetEmployees)
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello there")
	})
	e.GET("/employee/:id", controller.GetEmployeeByID)
	e.POST("/employee", controller.CreateEmployee)
	e.PUT("/employee", controller.UpdateEmployee)
	e.DELETE("/employee/:id", controller.DeleteEmployee)
}
func (ec Controller) GetEmployeeByID(ctx echo.Context) error {
	id := ctx.Param("id")

	employee, err := ec.eu.GetEmployeeByID(id)
	if err != nil {
		return ctx.JSON(501, model.ErrorResponse{
			Message: "Bad request",
		})
	}
	return ctx.JSON(200, employee)
}
func (ec Controller) GetEmployees(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")
	if page == "" || limit == "" {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Bad request",
		})
	}
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	employees, err := ec.eu.GetAllEmployees(pageInt, limitInt)

	if err != nil {
		log.Println(err)
		return ctx.JSON(501, model.ErrorResponse{
			Message: "Internal server error",
		})
	}
	return ctx.JSON(200, employees)
}
func (ec Controller) CreateEmployee(ctx echo.Context) error {
	request := new(model.Employee)
	if err := ctx.Bind(request); err != nil {
		log.Println("error here : ", err)
		return ctx.JSON(504, model.ErrorResponse{
			Message: "Bad request",
		})
	}
	err := ec.eu.CreateEmployee(request)
	if err != nil {
		log.Println(err)
		ctx.JSON(501, model.ErrorResponse{
			Message: "Internal server error",
		})
	}
	return ctx.JSON(201, "created")
}
func (ec Controller) UpdateEmployee(ctx echo.Context) error {
	request := new(model.Employee)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(504, model.ErrorResponse{
			Message: "Bad request",
		})
	}
	err := ec.eu.UpdateEmployee(request)
	if err != nil {
		log.Println(err)
		ctx.JSON(501, model.ErrorResponse{
			Message: "Internal server error",
		})
	}
	return nil
}
func (ec Controller) DeleteEmployee(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(504, model.ErrorResponse{
			Message: "Bad request.",
		})
	}
	err := ec.eu.DeleteEmployee(id)
	if err != nil {
		return ctx.JSON(501, model.ErrorResponse{
			Message: "Internal server error.",
		})
	}
	return ctx.JSON(200, "success")
}
