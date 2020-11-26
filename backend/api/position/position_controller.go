package position

import (
	"backend/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

type PositionController struct {
	positionUsecase PositionUsecase
}

func newPositionController(db *pg.DB) *PositionController {
	return &PositionController{positionUsecase: newPositionUsecase(db)}
}
func (pc PositionController) getAllPositions(ctx *gin.Context) {
	positions, err := pc.positionUsecase.getAllPosition()
	if err != nil {
		ctx.JSON(5004, "internal server error")
		return
	}

	ctx.JSON(200, gin.H{
		"code":   200,
		"status": "ok",
		"result": positions,
	})
}
func (pc PositionController) createPosition(ctx *gin.Context) {
	var pos Position
	err := ctx.BindJSON(&pos)
	if err != nil {
		ctx.JSON(504, "Failed parsing the JSON")
		return
	}
	err = helper.ValidateRequest(pos)
	if err != nil {
		ctx.JSON(400, "cannot be empty")
		return
	}
	err = pc.positionUsecase.createPosition(&pos)
	if err != nil {
		ctx.JSON(504, "Internal server error")
		return
	}
	ctx.JSON(200, gin.H{
		"code":   200,
		"status": "ok",
	})
}
func (pc PositionController) updatePosition(ctx *gin.Context) {
	var pos Position
	err := ctx.BindJSON(&pos)
	if err != nil {
		ctx.JSON(504, "Failed parsing the JSON")
		return
	}
	err = helper.ValidateRequest(pos)
	if err != nil {
		ctx.JSON(400, "cannot be empty")
		return
	}
	err = pc.positionUsecase.updatePosition(&pos)
	if err != nil {
		ctx.JSON(504, "Internal server error")
		return
	}
	ctx.JSON(200, gin.H{
		"code":   200,
		"status": "ok",
	})
}
func (pc PositionController) deletePosition(ctx *gin.Context) {
	id := ctx.Param("id")

	err := pc.positionUsecase.deletePosition(&id)
	if err != nil {
		ctx.JSON(504, "Internal server error")
		return
	}
	ctx.JSON(200, gin.H{
		"code":   200,
		"status": "ok",
	})
}
