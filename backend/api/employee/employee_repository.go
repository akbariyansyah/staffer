package employee

import (
	"github.com/go-pg/pg"
	"log"
)

type EmployeeRepositoryInterface interface {
	getAllEmployees() (*[]Employee, error)
	createEmployee(employee *Employee) (*Employee, error)
	updateEmployee(employee *Employee) error
	deleteEmployee(id *int) error
}
type EmployeeRepository struct {
	db *pg.DB
}

func (e EmployeeRepository) getAllEmployees() (*[]Employee, error) {
	var employees []Employee

	_, err := e.db.Query(&employees, "SELECT * FROM m_employee where is_delete = ? ", 0)
	if err != nil {
		panic(err)
		return nil, err
	}
	return &employees, nil
}

func (e EmployeeRepository) createEmployee(employee *Employee) (*int, error) {
	tx, err := e.db.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare(`INSERT INTO m_employee (id,full_name,birth_date,position_id,id_number,gender) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	res, err := stmt.Exec(&employee.ID, &employee.FullName, &employee.BirthDate, &employee.PositionID, &employee.IDNumber, &employee.Gender)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	userCreated := res.RowsReturned()
	tx.Commit()

	return &userCreated, nil
}

func (e EmployeeRepository) updateEmployee(emp *Employee) error {
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

func (e EmployeeRepository) deleteEmployee(id *string) error {
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

func newEmployeeRepository(db *pg.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}
