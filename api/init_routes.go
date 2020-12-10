package api

import (
	"database/sql"
	"github.com/labstack/echo"
	"staffer/api/employee"
)

func NewRoutes(e *echo.Echo, db *sql.DB) {
	employeeRepo := employee.NewEmployeeRepository(db)
	employeeUsecase := employee.NewEmployeeUsecase(employeeRepo)
	employee.NewEmployeeController(e, employeeUsecase)
}
