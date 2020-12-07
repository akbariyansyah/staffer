package config

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"staffer/api/employee"
)

func NewRoutes(e *echo.Echo, db *sql.DB) {
	employeeRepo := employee.NewEmployeeRepository(db)
	employeeUsecase := employee.NewEmployeeUsecase(employeeRepo)
	employee.NewEmployeeController(e, employeeUsecase)
}
