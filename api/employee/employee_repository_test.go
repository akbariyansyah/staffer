package employee

import (
	"database/sql"
	"log"
	"regexp"
	"staffer/model"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var employeeMock = &model.Employee{
	ID:        91,
	FullName:  "Keir Kreutzer",
	Email:     "kkreutzer0@networksolutions.com",
	Title:     "Mechanical Systems Engineer",
	Gender:    "M",
	Phone:     "(128) 3181004",
	Address:   "509 American Ash Avenue",
	IsMarried: true,
	BirthDate: time.Now(),
}

func TestNewEmployeeRepository(t *testing.T) {
	db, _ := NewRepoMock()
	empRepo := NewEmployeeRepository(db)
	assert.NotNil(t, empRepo)
	assert.NotEmpty(t, empRepo)
}

// NewMock -> initialize new mock of database.
func NewRepoMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
func TestNewRepoMock(t *testing.T) {
	db, mock := NewRepoMock()
	assert.NotNil(t, db)
	assert.NotEmpty(t, mock)
}
func TestEmployeeRepository_CountEmployees(t *testing.T) {
	db, mock := NewRepoMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	rows := mock.NewRows([]string{"total_data"}).AddRow(1000)

	mock.ExpectPrepare(regexp.QuoteMeta("select count(*) as total_data from employee")).ExpectQuery().WillReturnRows(rows)
	totalData, err := repo.CountEmployees()
	assert.NoError(t, err)
	assert.Equal(t, 1000, totalData)
	assert.NotNil(t, totalData)
}
func TestEmployeeRepository_CountEmployeesFail(t *testing.T) {
	db, mock := NewRepoMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	rows := mock.NewRows([]string{"total_data"})

	mock.ExpectPrepare(regexp.QuoteMeta("select count(*) as total_data from employee")).ExpectQuery().WillReturnRows(rows)
	totalData, err := repo.CountEmployees()
	assert.Error(t, err)
	assert.NotEqual(t, 1000, totalData)

}
func TestEmployeeRepository_GetEmployeeByID(t *testing.T) {
	db, mock := NewRepoMock()
	defer db.Close()
	repo := NewEmployeeRepository(db)
	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"}).AddRow(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate)
	mock.ExpectQuery(regexp.QuoteMeta("select * from employee where id = ?")).WithArgs("91").WillReturnRows(rows)

	employee, err := repo.GetEmployeeByID("91")
	assert.NoError(t, err)
	assert.NotNil(t, employee)
}
func TestEmployeeRepository_GetEmployeeByIDFail(t *testing.T) {
	db, mock := NewRepoMock()
	defer db.Close()
	repo := NewEmployeeRepository(db)
	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"})
	mock.ExpectQuery(regexp.QuoteMeta("select * from employee where id = ?")).WithArgs("90001").WillReturnRows(rows)

	employee, err := repo.GetEmployeeByID("90001")
	assert.Error(t, err)
	assert.NotNil(t, err)
	assert.Nil(t, employee)
}
func TestEmployeeRepository_GetEmployees(t *testing.T) {
	db, mock := NewRepoMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"}).AddRow(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate)

	mock.ExpectQuery(regexp.QuoteMeta("select * from employee limit ?,?")).WithArgs(1, 1).WillReturnRows(rows)

	employee, err := repo.GetAllEmployees(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, employee)
	assert.Len(t, employee, 1)
}
func TestEmployeeRepository_GetEmployeesFail(t *testing.T) {
	db, mock := NewRepoMock()
	repo := NewEmployeeRepository(db)

	defer db.Close()
	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"})

	mock.ExpectQuery(regexp.QuoteMeta("select * from employee limit ?,?")).WithArgs(0, 10).WillReturnRows(rows)

	employees, err := repo.GetAllEmployees(0, 10)
	assert.Error(t, err)
	assert.Nil(t, employees)
}
func TestEmployeeRepository_CreateEmployee(t *testing.T) {
	db, mock := NewRepoMock()

	repo := NewEmployeeRepository(db)

	repo.CreateEmployee(employeeMock)
	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("insert into employee(id,full_name,email,title,gender,phone,address,is_married,birth_date) values(?,?,?,?,?,?,?,?,?)")).ExpectExec().WithArgs(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.CreateEmployee(employeeMock)

	assert.NoError(t, err)
}
func TestEmployeeRepository_CreateEmployeeFail(t *testing.T) {
	db, mock := NewRepoMock()

	repo := NewEmployeeRepository(db)

	repo.CreateEmployee(employeeMock)
	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("insert into employee(id,full_name,email,title,gender,phone,address,is_married,birth_date) values(?,?,?,?,?,?,?,?,?)")).ExpectExec().WithArgs(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.CreateEmployee(employeeMock)

	assert.Error(t, err)
}
func TestEmployeeRepository_UpdateEmployee(t *testing.T) {
	db, mock := NewRepoMock()

	repo := NewEmployeeRepository(db)

	repo.CreateEmployee(employeeMock)
	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("update employee set full_name=?,email=?,title=?,gender=?,phone=?,address=?,is_married=?,birth_date=? where id=?")).ExpectExec().WithArgs(employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate, employeeMock.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.UpdateEmployee(employeeMock)

	assert.NoError(t, err)
}
func TestEmployeeRepository_UpdateEmployeeFail(t *testing.T) {
	db, mock := NewRepoMock()

	repo := NewEmployeeRepository(db)

	repo.CreateEmployee(employeeMock)
	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta("update employee set full_name=?,email=?,title=?,gender=?,phone=?,address=?,is_married=?,birth_date=? where id=?")).ExpectExec().WithArgs(employeeMock.FullName, employeeMock.Email, employeeMock.Title, employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate, employeeMock.ID).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.UpdateEmployee(employeeMock)

	assert.Error(t, err)
}
func TestEmployeeRepository_DeleteEmployee(t *testing.T) {
	db, mock := NewRepoMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("delete from employee where id = ?").ExpectExec().WithArgs(strconv.Itoa(employeeMock.ID)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.DeleteEmployee(strconv.Itoa(employeeMock.ID))
	assert.NoError(t, err)
	assert.Nil(t, err)
}
func TestEmployeeRepository_DeleteEmployeeFail(t *testing.T) {
	db, mock := NewRepoMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("delete from employee where id = ?").ExpectExec().WithArgs(strconv.Itoa(employeeMock.ID)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.DeleteEmployee(strconv.Itoa(employeeMock.ID))
	assert.Error(t, err)
	assert.NotNil(t, err)
}
