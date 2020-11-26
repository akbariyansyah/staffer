package employee

import (
	"github.com/go-pg/pg"
	"log"
)

type EmployeeRepository interface {
	getAllEmployees() (*[]Employee, error)
	createEmployee(employee *Employee) error
	updateEmployee(employee *Employee) error
	deleteEmployee(id *string) error
}

//  select me.id,me.full_name,me.birth_date,mp.name,me.id_number,me.gender from m_employee me join m_position mp on (me.position_id=mp.id)
type EmployeeRepositoryImpl struct {
	db *pg.DB
}

func (e EmployeeRepositoryImpl) getAllEmployees() (*[]Employee, error) {
	var employees []Employee

	_, err := e.db.Query(&employees, "SELECT * FROM m_employee where is_delete = ? ", 0)
	if err != nil {
		panic(err)
		return nil, err
	}
	return &employees, nil
}

func (e EmployeeRepositoryImpl) createEmployee(employee *Employee) error {
	tx, err := e.db.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare(`INSERT INTO m_employee (id,full_name,birth_date,position_id,id_number,gender) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(&employee.ID, &employee.FullName, &employee.BirthDate, &employee.PositionID, &employee.IDNumber, &employee.Gender)
	if err != nil {
		tx.Rollback()
		log.Println(err)

		return err
	}
	//userCreated := res.RowsReturned()
	tx.Commit()
	log.Println(err)
	return nil
}

func (e EmployeeRepositoryImpl) updateEmployee(emp *Employee) error {
	tx, err := e.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`UPDATE m_employee set full_name = $1 ,birth_date = $2,position_id = $3 ,id_number = $4 ,gender = $5 WHERE id = $6`)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(emp.FullName, emp.BirthDate, emp.PositionID, emp.IDNumber, emp.Gender, emp.ID)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (e EmployeeRepositoryImpl) deleteEmployee(id *string) error {
	tx, err := e.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	stmt, err := tx.Prepare(`UPDATE m_employee set is_delete = $1 where id = $2`)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(1, *id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()

	return nil
}

func newEmployeeRepository(db *pg.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}
