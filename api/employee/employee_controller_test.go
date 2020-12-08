package employee

import (
	"net/http"
	"net/http/httptest"
	"staffer/config"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func NewControllerMock() *echo.Echo {
	conf := config.NewConfig()
	db := config.NewDatabase(conf)
	empRepo := NewEmployeeRepository(db)
	empUsecase := NewEmployeeUsecase(empRepo)
	controller := Controller{empUsecase}

	e := echo.New()

	e.GET("/employee", controller.GetEmployees)

	return e
}
func TestController_GetEmployees(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:4000/employee?page=1&limit=10", nil)
	if err != nil {
		t.Fatalf("Failed : %v", err)
	}
	resp := httptest.NewRecorder()
	NewControllerMock().ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Result().StatusCode, "Response is expected")
}
func TestController_GetEmployeesFail(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:4000/employee?page=0&limit=10", nil)
	if err != nil {
		t.Fatalf("Failed : %v", err)
	}
	resp := httptest.NewRecorder()
	NewControllerMock().ServeHTTP(resp, req)
	assert.Equal(t, 501, resp.Result().StatusCode, "Response is expected")
}
