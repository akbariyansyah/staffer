package employee

import "github.com/go-pg/pg"

type EmployeeUsecaseInterface interface {
	getAllEmployees() (*[]Employee, error)
	createEmployee(employee *Employee) error
	updateEmployee(employee *Employee) error
	deleteEmployee(id *string) error
}
type EmployeeUsecase struct {
	employeeRepo *EmployeeRepository
}

func (e EmployeeUsecase) getAllEmployees() (*[]Employee, error) {
	result, err := e.employeeRepo.getAllEmployees()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e EmployeeUsecase) createEmployee(employee *Employee) error {
	panic("implement me")
}

func (e EmployeeUsecase) updateEmployee(employee *Employee) error {
	panic("implement me")
}

func (e EmployeeUsecase) deleteEmployee(id *string) error {
	panic("implement me")
}

func newEmployeeUsecase(db *pg.DB) *EmployeeUsecase {
	return &EmployeeUsecase{employeeRepo: newEmployeeRepository(db)}
}
