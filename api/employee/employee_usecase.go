package employee

import (
	"errors"
	"staffer/model"
)

type IEmployeeUsecase interface {
	GetAllEmployees(page, limit int) (map[string]interface{}, error)
	UpdateEmployee(emp *model.Employee) error
	CreateEmployee(emp *model.Employee) error
	DeleteEmployee(id string) error
}
type EmployeeUsecase struct {
	empRepo IEmployeeRepository
}

func NewEmployeeUsecase(empRepo IEmployeeRepository) IEmployeeUsecase {
	return &EmployeeUsecase{empRepo: empRepo}
}
func (eu EmployeeUsecase) GetAllEmployees(page, limit int) (map[string]interface{}, error) {

	offset := (page * limit) - limit
	totalData, err := eu.empRepo.CountEmployees()

	if err != nil {
		return nil, err
	}
	if totalData <= offset {
		return nil, errors.New("Requested page is not available")
	}
	emp, err := eu.empRepo.GetAllEmployees(offset, limit)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"totalData": totalData,
		"data":      emp,
	}, nil
}
func (eu EmployeeUsecase) UpdateEmployee(emp *model.Employee) error {
	err := eu.empRepo.UpdateEmployee(emp)
	if err != nil {
		return err
	}
	return nil
}
func (eu EmployeeUsecase) CreateEmployee(emp *model.Employee) error {
	err := eu.empRepo.CreateEmployee(emp)
	if err != nil {
		return err
	}
	return nil
}
func (eu EmployeeUsecase) DeleteEmployee(id string) error {
	err := eu.empRepo.DeleteEmployee(id)
	if err != nil {
		return err
	}
	return nil
}
