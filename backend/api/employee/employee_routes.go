package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func NewEmployeeRoutes(route string, g *gin.Engine, pg *pg.DB) {
	g.GET(route, getAllEmployees)
	g.POST(route, createEmployee)
	g.PUT(route, updateEmployee)
	g.DELETE(route, deleteEmployee)
}
func getAllEmployees(ctx *gin.Context) {
	ctx.JSON(200, "its work")
}
func createEmployee(ctx *gin.Context) {
	ctx.JSON(200, "its work")
}
func updateEmployee(ctx *gin.Context) {
	ctx.JSON(200, "its work")
}
func deleteEmployee(ctx *gin.Context) {
	ctx.JSON(200, "its work")
}
