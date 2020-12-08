package employee

import (
	"database/sql"
	"log"
	"regexp"
	"staffer/model"
	"testing"

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
	BirthDate: "1987-06-13T00:00:00Z",
}

// NewMock -> initialize new mock of database.
func NewRepoMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
func TestRepo_CountEmployees(t *testing.T) {
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
func TestRepo_CountEmployeesFail(t *testing.T) {
	db, mock := NewRepoMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	rows := mock.NewRows([]string{"total_data"})

	mock.ExpectPrepare(regexp.QuoteMeta("select count(*) as total_data from employee")).ExpectQuery().WillReturnRows(rows)
	totalData, err := repo.CountEmployees()
	assert.Error(t, err)
	assert.NotEqual(t, 1000, totalData)

}
func TestRepo_GetEmployees(t *testing.T) {
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
func TestRepo_GetEmployeesFail(t *testing.T) {
	db, mock := NewRepoMock()
	repo := NewEmployeeRepository(db)

	defer db.Close()
	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"})

	mock.ExpectQuery(regexp.QuoteMeta("select * from employee limit ?,?")).WithArgs(0, 10).WillReturnRows(rows)

	employees, err := repo.GetAllEmployees(0, 10)
	assert.Error(t, err)
	assert.Nil(t, employees)
}
