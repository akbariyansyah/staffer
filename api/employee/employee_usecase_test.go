package employee

import (
	"staffer/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newUsecaseMock() IEmployeeUsecase {
	conf := config.NewConfig()
	db,err := config.NewDatabase(conf)
	if err != nil {
		panic(err)
	}
	empRepo := NewEmployeeRepository(db)
	empUsecase := NewEmployeeUsecase(empRepo)
	return empUsecase
}
func TestUsecase_GetAllEmployees(t *testing.T) {
	employees, err := newUsecaseMock().GetAllEmployees(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, employees)
	assert.Len(t, employees, 2)
}
func TestUsecase_GetAllEmployeesFail(t *testing.T) {
	employees, err := newUsecaseMock().GetAllEmployees(0, 1)
	assert.Error(t, err)
	assert.Nil(t, employees)
}
