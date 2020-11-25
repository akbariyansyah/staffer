package position

import (
	"github.com/go-pg/pg"
)

type PositionUsecase interface {
	getAllPosition() (*[]Position, error)
	createPosition(pos *Position) error
	updatePosition(pos *Position) error
	deletePosition(id *string) error
}
type PositionUsecaseImpl struct {
	positionRepo PositionRepository
}

func (p PositionUsecaseImpl) getAllPosition() (*[]Position, error) {
	positions, err := p.positionRepo.getAllPositions()
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (p PositionUsecaseImpl) createPosition(pos *Position) error {
	err := p.positionRepo.createPosition(pos)
	if err != nil {
		return err
	}
	return nil
}

func (p PositionUsecaseImpl) updatePosition(pos *Position) error {
	err := p.positionRepo.updatePosition(pos)
	if err != nil {
		return err
	}
	return nil
}

func (p PositionUsecaseImpl) deletePosition(id *string) error {
	err := p.positionRepo.deletePosition(id)
	if err != nil {
		return err
	}
	return nil
}

func newPositionUsecase(db *pg.DB) PositionUsecase {
	return &PositionUsecaseImpl{positionRepo: newPositionRepository(db)}
}
