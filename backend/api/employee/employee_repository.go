package employee

import "github.com/go-pg/pg"

type EmployeeRepositoryInterface interface {
	getAllEmployees() (*[]Employee, error)
	createEmployee(employee *Employee) error
	updateEmployee(employee *Employee) error
	deleteEmployee(id *int) error
}
type EmployeeRepository struct {
	db           *pg.DB
	employeeRepo EmployeeRepositoryInterface
}

func (e EmployeeRepository) getAllEmployees() (*[]Employee, error) {
	panic("implement me")
}

func (e EmployeeRepository) createEmployee(employee *Employee) error {
	panic("implement me")
}

func (e EmployeeRepository) updateEmployee(employee *Employee) error {
	panic("implement me")
}

func (e EmployeeRepository) deleteEmployee(id *int) error {
	panic("implement me")
}

func newEmployeeRepository(db *pg.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}
