package employee

import (
	"database/sql"
	"errors"
	"log"
	"staffer/model"
)

// IEmployeeRepository -> available method on this layer.
type IEmployeeRepository interface {
	GetAllEmployees(offset, limit int) (*[]model.Employee, error)
	CountEmployees() (int, error)
}

// EmployeeRepository -> type that contains the database connection.
type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(DB *sql.DB) IEmployeeRepository {
	return &EmployeeRepository{DB: DB}
}

// GetAllEmployees -> retrieve all employee from database.
func (db EmployeeRepository) GetAllEmployees(offset, limit int) (*[]model.Employee, error) {

	var employees = []model.Employee{}

	rows, err := db.DB.Query("select * from employee limit ?,?", offset, limit)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		emp := new(model.Employee)
		err := rows.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Title, &emp.Gender, &emp.Phone, &emp.Address, &emp.IsMarried, &emp.BirthDate)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		employees = append(employees, *emp)
	}
	if len(employees) == 0 {
		log.Println("no employee")
		return nil, errors.New("Offset doesn't found")
	}
	return &employees, nil
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
