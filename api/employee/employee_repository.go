package employee

import (
	"database/sql"
	"errors"
	"log"
	"staffer/model"
)

// IEmployeeRepository -> available method on this layer.
type IEmployeeRepository interface {
	GetAllEmployees(offset, limit int) ([]*model.Employee, error)
	GetEmployeeByID(id string) (*model.Employee, error)
	CountEmployees() (int, error)
	UpdateEmployee(*model.Employee) error
	CreateEmployee(*model.Employee) error
	DeleteEmployee(id string) error
}

// EmployeeRepository -> type that contains the database connection.
type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(DB *sql.DB) IEmployeeRepository {
	return &EmployeeRepository{DB: DB}
}
func (db EmployeeRepository) GetEmployeeByID(id string) (*model.Employee, error) {
	employee := new(model.Employee)

	err := db.DB.QueryRow("select * from employee where id = ?", id).Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Title, &employee.Gender, &employee.Phone, &employee.Address, &employee.IsMarried, &employee.BirthDate)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return employee, nil
}

// GetAllEmployees -> retrieve all employee from database.
func (db EmployeeRepository) GetAllEmployees(offset, limit int) ([]*model.Employee, error) {

	var employees = []*model.Employee{}

	rows, err := db.DB.Query("select * from employee limit ?,?", offset, limit)

	if err != nil {
		log.Printf("FAILED : %v", err)
		return nil, err
	}
	for rows.Next() {
		emp := new(model.Employee)
		err := rows.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Title, &emp.Gender, &emp.Phone, &emp.Address, &emp.IsMarried, &emp.BirthDate)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		employees = append(employees, emp)
	}
	if len(employees) == 0 {
		return nil, errors.New("Offset doesn't found")
	}
	return employees, nil
}

// CountEmployees -> count all the employee data from database.
func (db EmployeeRepository) CountEmployees() (int, error) {
	var totalData int
	stmt, err := db.DB.Prepare("select count(*) as total_data from employee")
	if err != nil {
		return totalData, err
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&totalData)
	if err != nil {
		return totalData, err
	}
	return totalData, nil
}
func (db EmployeeRepository) UpdateEmployee(emp *model.Employee) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update employee set full_name=?,email=?,title=?,gender=?,phone=?,address=?,is_married=?,birth_date=? where id = ?")

	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(emp.FullName, emp.Email, emp.Title, emp.Gender, emp.Phone, emp.Address, emp.IsMarried, emp.BirthDate, emp.ID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	return tx.Commit()
}
func (db EmployeeRepository) CreateEmployee(emp *model.Employee) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into employee(id,full_name,email,title,gender,phone,address,is_married,birth_date) values(?,?,?,?,?,?,?,?,?)")

	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(emp.ID, emp.FullName, emp.Email, emp.Title, emp.Gender, emp.Phone, emp.Address, emp.IsMarried, emp.BirthDate)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	return tx.Commit()
}
func (db EmployeeRepository) DeleteEmployee(id string) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("delete from employee where id = ?")

	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	return tx.Commit()
}
