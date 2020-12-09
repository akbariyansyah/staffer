package api

import (
	"database/sql"
	"staffer/api/employee"

	"github.com/labstack/echo"
)

func NewRoutes(e *echo.Echo, db *sql.DB) {
	employeeRepo := employee.NewEmployeeRepository(db)
	employeeUsecase := employee.NewEmployeeUsecase(employeeRepo)
	employee.NewEmployeeController(e, employeeUsecase)
}
