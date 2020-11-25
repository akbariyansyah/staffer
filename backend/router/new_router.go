package router

import (
	"backend/api/employee"
	"backend/api/position"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

const (
	employeeRoutes = "/employee"
	positionRoutes = "/position"
)

func NewRouter(g *gin.Engine, pg *pg.DB) {
	employee.NewEmployeeRoutes(employeeRoutes, g, pg)
	position.NewPositionRoutes(positionRoutes, g, pg)
}
