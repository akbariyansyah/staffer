package employee

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"staffer/config"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func NewControllerMock() *echo.Echo {
	conf := config.NewConfig()
	db, err := config.NewDatabase(conf)
	if err != nil {
		panic(err)
	}
	empRepo := NewEmployeeRepository(db)
	empUsecase := NewEmployeeUsecase(empRepo)
	controller := Controller{empUsecase}

	e := echo.New()

	e.GET("/employee", controller.GetEmployees)

	return e
}

type Pagination struct {
	name  string
	page  string
	limit string
}

func TestController_GetEmployees(t *testing.T) {
	tt := []Pagination{
		{
			name:  "test 1",
			page:  "10",
			limit: "5",
		},
		{
			name:  "test 2",
			page:  "100",
			limit: "5",
		},
		{
			name:  "test 3",
			page:  "10",
			limit: "50",
		},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			endpoint := fmt.Sprintf("http://localhost:4000/employee?page=%s&limit=%s", v.page, v.limit)
			req, err := http.NewRequest("GET", endpoint, nil)
			if err != nil {
				t.Fatalf("Failed : %v", err)
			}
			resp := httptest.NewRecorder()
			NewControllerMock().ServeHTTP(resp, req)
			assert.Equal(t, 200, resp.Result().StatusCode, "Response is expected")
		})
	}
}

func TestController_GetEmployeesFail(t *testing.T) {
	tt := []Pagination{
		{
			name:  "test 1",
			page:  "0",
			limit: "-5",
		},
		{
			name:  "test 2",
			page:  "-6",
			limit: "0",
		},
		{
			name:  "test 3",
			page:  "-0",
			limit: "-10",
		},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			endpoint := fmt.Sprintf("http://localhost:4000/employee?page=%s&limit=%s", v.page, v.limit)
			req, err := http.NewRequest("GET", endpoint, nil)
			if err != nil {
				t.Fatalf("Failed : %v", err)
			}
			resp := httptest.NewRecorder()
			NewControllerMock().ServeHTTP(resp, req)
			assert.Equal(t, 501, resp.Result().StatusCode, "Response is expected")
		})
	}
}
