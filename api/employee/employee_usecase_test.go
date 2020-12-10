package employee

import (
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newUsecaseMock() (IEmployeeUsecase, sqlmock.Sqlmock) {
	db, mock := NewRepoMock()
	empRepo := NewEmployeeRepository(db)
	empUsecase := NewEmployeeUsecase(empRepo)
	return empUsecase, mock
}

func TestUsecase_GetAllEmployees(t *testing.T) {
	empusecase, mock := newUsecaseMock()

	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"}).AddRow(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate)
	rows2 := mock.NewRows([]string{"total_data"}).AddRow(1000)
	mock.ExpectPrepare(regexp.QuoteMeta("select count(*) as total_data from employee")).ExpectQuery().WillReturnRows(rows2)
	mock.ExpectQuery(regexp.QuoteMeta("select * from employee limit ?,?")).WithArgs(0, 1).WillReturnRows(rows)

	employee, err := empusecase.GetAllEmployees(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, employee)
}
func TestUsecase_GetAllEmployeesFail(t *testing.T) {
	empusecase, mock := newUsecaseMock()
	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"})

	mock.ExpectQuery(regexp.QuoteMeta("select * from employee limit ?,?")).WithArgs(0, 10).WillReturnRows(rows)
	employees, err := empusecase.GetAllEmployees(0, 10)
	assert.Error(t, err)
	assert.Nil(t, employees)
}
func TestEmployeeUsecase_GetEmployeeByID(t *testing.T) {
	empUsecase, mock := newUsecaseMock()
	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"}).AddRow(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate)
	mock.ExpectQuery(regexp.QuoteMeta("select * from employee where id = ?")).WithArgs("1").WillReturnRows(rows)
	emp, err := empUsecase.GetEmployeeByID("1")
	assert.NoError(t, err)
	assert.NotNil(t, emp)
}

func TestEmployeeUsecase_GetEmployeeByIDFail(t *testing.T) {
	empUsecase, mock := newUsecaseMock()
	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"})
	mock.ExpectQuery(regexp.QuoteMeta("select * from employee where id = ?")).WithArgs("1000").WillReturnRows(rows)
	emp, err := empUsecase.GetEmployeeByID("1000")
	assert.Error(t, err)
	assert.Nil(t, emp)
}
func TestEmployeeUsecase_CreateEmployee(t *testing.T) {
	empUsecase, mock := newUsecaseMock()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("insert into employee(id,full_name,email,title,gender,phone,address,is_married,birth_date) values(?,?,?,?,?,?,?,?,?)")).ExpectExec().WithArgs(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := empUsecase.CreateEmployee(employeeMock)

	assert.NoError(t, err)
}
func TestEmployeeUsecase_CreateEmployeeFail(t *testing.T) {
	empUsecase, mock := newUsecaseMock()
	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("insert into employee(id,full_name,email,title,gender,phone,address,is_married,birth_date) values(?,?,?,?,?,?,?,?,?)")).ExpectExec().WithArgs(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := empUsecase.CreateEmployee(employeeMock)

	assert.Error(t, err)
}
func TestEmployeeUsecase_UpdateEmployee(t *testing.T) {
	empUsecase, mock := newUsecaseMock()
	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("update employee set full_name=?,email=?,title=?,gender=?,phone=?,address=?,is_married=?,birth_date=? where id=?")).ExpectExec().WithArgs(employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate, employeeMock.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := empUsecase.UpdateEmployee(employeeMock)

	assert.NoError(t, err)
}
func TestEmployeeUsecase_UpdateEmployeeFail(t *testing.T) {
	empUsecase, mock := newUsecaseMock()

	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("update employee set full_name=?,email=?,title=?,gender=?,phone=?,address=?,is_married=?,birth_date=? where id=?")).ExpectExec().WithArgs(employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate, employeeMock.ID).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := empUsecase.UpdateEmployee(employeeMock)

	assert.Error(t, err)
}
func TestEmployeeUsecase_DeleteEmployee(t *testing.T) {
	empUsecase, mock := newUsecaseMock()

	mock.ExpectBegin()
	mock.ExpectPrepare("delete from employee where id = ?").ExpectExec().WithArgs(strconv.Itoa(employeeMock.ID)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := empUsecase.DeleteEmployee(strconv.Itoa(employeeMock.ID))
	assert.NoError(t, err)
	assert.Nil(t, err)
}
func TestEmployeeUsecase_DeleteEmployeeFail(t *testing.T) {
	empUsecase, mock := newUsecaseMock()
	mock.ExpectBegin()
	mock.ExpectPrepare("delete from employee where id = ?").ExpectExec().WithArgs(strconv.Itoa(employeeMock.ID)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := empUsecase.DeleteEmployee(strconv.Itoa(employeeMock.ID))
	assert.Error(t, err)
	assert.NotNil(t, err)
}
