package employee

import (
	"staffer/model"
)

type IEmployeeUsecase interface {
	GetAllEmployees() (*[]model.Employee, error)
}
type EmployeeUsecase struct {
	empRepo IEmployeeRepository
}

func NewEmployeeUsecase(empRepo IEmployeeRepository) IEmployeeUsecase {
	return &EmployeeUsecase{empRepo: empRepo}
}
func (eu EmployeeUsecase) GetAllEmployees() (*[]model.Employee, error) {
	emp, err := eu.empRepo.GetAllEmployees()
	if err != nil {
		return nil, err
	}
	return emp, nil
}
