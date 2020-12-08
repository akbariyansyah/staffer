package employee

import (
	"errors"
	"strconv"
)

type IEmployeeUsecase interface {
	GetAllEmployees(page, limit string) (map[string]interface{}, error)
}
type EmployeeUsecase struct {
	empRepo IEmployeeRepository
}

func NewEmployeeUsecase(empRepo IEmployeeRepository) IEmployeeUsecase {
	return &EmployeeUsecase{empRepo: empRepo}
}
func (eu EmployeeUsecase) GetAllEmployees(page, limit string) (map[string]interface{}, error) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt * limitInt) - limitInt
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
