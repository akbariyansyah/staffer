package position

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func NewPositionRoutes(route string, g *gin.Engine, pg *pg.DB) {
	positionController := newPositionController(pg)
	g.GET(route, positionController.getAllPositions)
	g.POST(route, positionController.createPosition)
	g.PUT(route, positionController.updatePosition)
	g.DELETE(route+"/:id", positionController.deletePosition)

}
