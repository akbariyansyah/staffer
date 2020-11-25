package employee

import "github.com/go-pg/pg"

type EmployeeUsecase interface {
	getAllEmployees() (*[]Employee, error)
	createEmployee(employee *Employee) error
	updateEmployee(employee *Employee) error
	deleteEmployee(id *string) error
}
type EmployeeUsecaseImpl struct {
	employeeRepo EmployeeRepository
}

func (e EmployeeUsecaseImpl) getAllEmployees() (*[]Employee, error) {
	result, err := e.employeeRepo.getAllEmployees()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e EmployeeUsecaseImpl) createEmployee(employee *Employee) error {
	err := e.employeeRepo.createEmployee(employee)
	if err != nil {
		return  err
	}
	return nil
}

func (e EmployeeUsecaseImpl) updateEmployee(employee *Employee) error {
	err := e.employeeRepo.updateEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeeUsecaseImpl) deleteEmployee(id *string) error {
	err := e.employeeRepo.deleteEmployee(id)
	if err != nil {
		return err
	}
	return nil
}

func newEmployeeUsecase(db *pg.DB) EmployeeUsecase {
	return &EmployeeUsecaseImpl{newEmployeeRepository(db)}
}
