package employee

import "github.com/go-pg/pg"

type EmployeeUsecaseInterface interface {
	getAllEmployees() (*[]Employee, error)
	createEmployee(employee *Employee) (*int,error)
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

func (e EmployeeUsecase) createEmployee(employee *Employee) (*int,error) {
	result,err := e.employeeRepo.createEmployee(employee)
	if err != nil {
		return nil, err
	}
	return result,nil
}

func (e EmployeeUsecase) updateEmployee(employee *Employee) error {
	err := e.employeeRepo.updateEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeeUsecase) deleteEmployee(id *string) error {
	err := e.employeeRepo.deleteEmployee(id)
	if err != nil {
		return err
	}
	return nil
}

func newEmployeeUsecase(db *pg.DB) *EmployeeUsecase {
	return &EmployeeUsecase{employeeRepo: newEmployeeRepository(db)}
}
