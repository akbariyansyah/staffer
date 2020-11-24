package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"backend/api/employee"
)

const (
	employeeRoutes = "/employee"
	positionRoutes = "/position"
)

func NewRouter(g *gin.Engine, pg *pg.DB) {
	employee.NewEmployeeRoutes(employeeRoutes,g,pg)
}
