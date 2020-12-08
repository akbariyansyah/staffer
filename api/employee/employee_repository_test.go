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
func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
func TestCountEmployees(t *testing.T) {
	db, mock := NewMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	rows := mock.NewRows([]string{"total_data"}).AddRow(1000)

	mock.ExpectPrepare(regexp.QuoteMeta("select count(*) as total_data from employee")).ExpectQuery().WillReturnRows(rows)
	totalData, err := repo.CountEmployees()
	assert.NoError(t, err)
	assert.Equal(t, 1000, totalData)
	assert.NotNil(t, totalData)
}
func TestCountEmployeesFail(t *testing.T) {
	db, mock := NewMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	rows := mock.NewRows([]string{"total_data"})

	mock.ExpectPrepare(regexp.QuoteMeta("select count(*) as total_data from employee")).ExpectQuery().WillReturnRows(rows)
	totalData, err := repo.CountEmployees()
	assert.Error(t, err)
	assert.NotEqual(t, 1000, totalData)

}
func TestGetEmployees(t *testing.T) {
	db, mock := NewMock()
	repo := NewEmployeeRepository(db)
	defer db.Close()

	rows := mock.NewRows([]string{"id", "full_name", "email", "title", "gender", "phone", "address", "is_married", "birth_date"}).AddRow(employeeMock.ID, employeeMock.FullName, employeeMock.Email, employeeMock.Title,employeeMock.Gender, employeeMock.Phone, employeeMock.Address, employeeMock.IsMarried, employeeMock.BirthDate)

	mock.ExpectQuery(regexp.QuoteMeta("select * from employee limit ?,?")).WithArgs(0, 1).WillReturnRows(rows)

	employee, err := repo.GetAllEmployees(0, 1)
	assert.NoError(t, err)
	assert.NotNil(t, employee)

}
