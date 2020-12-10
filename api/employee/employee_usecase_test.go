package employee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newUsecaseMock() IEmployeeUsecase {
	db, _ := NewRepoMock()
	empRepo := NewEmployeeRepository(db)
	empUsecase := NewEmployeeUsecase(empRepo)
	return empUsecase
}
func TestUsecase_GetAllEmployees(t *testing.T) {
	usecase := newUsecaseMock()
	employees, err := usecase.GetAllEmployees(1, 10)
	assert.NoError(t, err)
	assert.NotNil(t, employees)
	assert.Len(t, employees, 2)
}
func TestUsecase_GetAllEmployeesFail(t *testing.T) {
	employees, err := newUsecaseMock().GetAllEmployees(0, 1)
	assert.Error(t, err)
	assert.Nil(t, employees)
}
func TestEmployeeUsecase_GetEmployeeByID(t *testing.T) {
	employee, err := newUsecaseMock().GetEmployeeByID("1")
	assert.NoError(t, err)
	assert.NotNil(t, employee)
}
func TestEmployeeUsecase_GetEmployeeByIDFail(t *testing.T) {
	employee, err := newUsecaseMock().GetEmployeeByID("1000000")
	assert.Error(t, err)
	assert.Nil(t, employee)
}
func TestEmployeeUsecase_UpdateEmployee(t *testing.T) {
	err := newUsecaseMock().UpdateEmployee(employeeMock)
	assert.Nil(t, err)
}
