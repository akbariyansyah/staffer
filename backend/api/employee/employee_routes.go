package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func NewEmployeeRoutes(route string, g *gin.Engine, pg *pg.DB) {
	employeeController := newEmployeeController(pg)
	g.GET(route, employeeController.getAllEmployees)
	g.POST(route, employeeController.createEmployee)
	g.PUT(route, employeeController.updateEmployee)
	g.DELETE(route+"/:id", employeeController.deleteEmployee)
}
