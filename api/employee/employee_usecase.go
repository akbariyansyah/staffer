package employee

import (
	"errors"
)

type IEmployeeUsecase interface {
	GetAllEmployees(page, limit int) (map[string]interface{}, error)
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
