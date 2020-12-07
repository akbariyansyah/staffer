package employee

import (
	"database/sql"
	"log"
	"staffer/model"
)

type IEmployeeRepository interface {
	GetAllEmployees() (*[]model.Employee, error)
}
type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(DB *sql.DB) IEmployeeRepository {
	return &EmployeeRepository{DB: DB}
}
// GetAllEmployees -> retrieve all employee from database.
func (db EmployeeRepository) GetAllEmployees() (*[]model.Employee, error) {
	var employees = []model.Employee{}

	rows, err := db.DB.Query("select * from employee")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		emp := new(model.Employee)
		err := rows.Scan(&emp.ID, &emp.FullName,&emp.Email, &emp.Title, &emp.Gender, &emp.Phone, &emp.Address, &emp.IsMarried, &emp.BirthDate)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		employees = append(employees, *emp)
	}
	return &employees, nil
}
