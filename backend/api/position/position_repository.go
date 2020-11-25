package position

import (
	"github.com/go-pg/pg"
	"log"
)

type PositionRepository interface {
	getAllPositions() (*[]Position, error)
	createPosition(pos *Position) error
	updatePosition(pos *Position) error
	deletePosition(id *string) error
}

type PositionRepositoryImpl struct {
	db *pg.DB
}

func (p PositionRepositoryImpl) getAllPositions() (*[]Position, error) {
	var positions []Position

	_, err := p.db.Query(&positions, `SELECT * FROM m_position where is_delete = ?`, 0)
	if err != nil {
		log.Println("ERROR : ", err)
		return nil, err
	}
	return &positions, err
}

func (p PositionRepositoryImpl) createPosition(pos *Position) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("ERROR : ", err)
		return err
	}
	stmt, err := tx.Prepare(`INSERT INTO m_position (code,name) VALUES ($1,$2)`)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	_, err = stmt.Exec(&pos.Code, &pos.Name)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	tx.Commit()
	return nil

}

func (p PositionRepositoryImpl) updatePosition(pos *Position) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("ERROR : ", err)
		return err
	}
	stmt, err := tx.Prepare(`UPDATE m_position set code = $1 ,name = $2 where id = $3`)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	_, err = stmt.Exec(pos.Code, pos.Name, pos.ID)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	tx.Commit()
	return nil
}

func (p PositionRepositoryImpl) deletePosition(id *string) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("ERROR : ", err)
		return err
	}
	stmt, err := tx.Prepare(`UPDATE m_position set is_delete = $1 where id = $2`)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	_, err = stmt.Exec(1, *id)
	if err != nil {
		tx.Rollback()
		log.Println("ERROR : ", err)
		return err
	}
	tx.Commit()
	return nil

}

func newPositionRepository(db *pg.DB) PositionRepository {
	return &PositionRepositoryImpl{db: db}
}
