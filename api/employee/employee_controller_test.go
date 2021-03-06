package employee

import (
	"github.com/DATA-DOG/go-sqlmock"
	"staffer/model"
	"testing"
)

func NewControllerMock() Controller {

	db, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	empRepo := NewEmployeeRepository(db)
	empUsecase := NewEmployeeUsecase(empRepo)
	controller := Controller{eu: empUsecase}

	return controller
}

type Pagination struct {
	name  string
	page  string
	limit string
}

//func TestController_GetEmployees(t *testing.T) {
//	tt := []Pagination{
//		{
//			name:  "test 1",
//			page:  "10",
//			limit: "5",
//		},
//		{
//			name:  "test 2",
//			page:  "100",
//			limit: "5",
//		},
//		{
//			name:  "test 3",
//			page:  "10",
//			limit: "50",
//		},
//	}
//	for _, v := range tt {
//		t.Run(v.name, func(t *testing.T) {
//			endpoint := fmt.Sprintf("http://localhost:4000/employee?page=%s&limit=%s", v.page, v.limit)
//			req, err := http.NewRequest("GET", endpoint, nil)
//			if err != nil {
//				t.Fatalf("Failed : %v", err)
//			}
//			resp := httptest.NewRecorder()
//			NewControllerMock().ServeHTTP(resp, req)
//			assert.Equal(t, 200, resp.Result().StatusCode, "Response is expected")
//		})
//	}
//}
//
//func TestController_GetEmployeesFail(t *testing.T) {
//	tt := []Pagination{
//		{
//			name:  "test 1",
//			page:  "0",
//			limit: "-5",
//		},
//		{
//			name:  "test 2",
//			page:  "-6",
//			limit: "0",
//		},
//		{
//			name:  "test 3",
//			page:  "-0",
//			limit: "-10",
//		},
//	}
//	for _, v := range tt {
//		t.Run(v.name, func(t *testing.T) {
//			endpoint := fmt.Sprintf("http://localhost:4000/employee?page=%s&limit=%s", v.page, v.limit)
//			req, err := http.NewRequest("GET", endpoint, nil)
//			if err != nil {
//				t.Fatalf("Failed : %v", err)
//			}
//			resp := httptest.NewRecorder()
//			NewControllerMock().ServeHTTP(resp, req)
//			assert.Equal(t, 501, resp.Result().StatusCode, "Response is expected")
//		})
//	}
//}
//func TestController_GetEmployeeByIDFail(t *testing.T) {
//	req, err := http.NewRequest("GET", "http://localhost:4000/employee/9999", nil)
//	if err != nil {
//		t.Fatalf("Failed : %v", err)
//	}
//	resp := httptest.NewRecorder()
//	NewControllerMock().ServeHTTP(resp, req)
//	assert.Equal(t, `{"message":"Bad request"}
//`, resp.Body.String(), "response is expected")
//	assert.Equal(t, 501, resp.Result().StatusCode, "response is expected")
//
//}
//func TestController_GetEmployeeByID(t *testing.T) {
//	req, err := http.NewRequest("GET", "http://localhost:4000/employee/100", nil)
//	if err != nil {
//		t.Fatalf("Failed : %v", err)
//	}
//	resp := httptest.NewRecorder()
//	NewControllerMock().ServeHTTP(resp, req)
//	assert.Equal(t, 200, resp.Result().StatusCode, "Expected result found")
//}
type handler struct {
	db map[string]*model.Employee
}

//func TestController_CreateEmployee(t *testing.T) {
//	byteOfEmployee, _ := json.Marshal(repo.employeeMock)
//	req, err := http.NewRequest("POST", "/", bytes.NewReader(byteOfEmployee))
//	if err != nil {
//		t.Fatalf("Failed : %v", err)
//	}
//	req.Header.Set("Content-Type", "application/json")
//
//	resp := httptest.NewRecorder()
//
//	e := echo.New()
//
//	c := e.NewContext(req, resp)
//	c.SetPath("/amployee")
//
//	h := NewControllerMock()
//
//	if assert.NoError(t, h.CreateEmployee(c)) {
//		assert.Equal(t, http.StatusCreated, resp.Code, "expected code")
//	}
//}
//func TestController_CreateEmployeeFail(t *testing.T) {
//	byteOfEmployee, _ := json.Marshal(repo.employeeMock)
//	req, err := http.NewRequest("POST", "http://localhost:4000/employee", bytes.NewBuffer(byteOfEmployee))
//	if err != nil {
//		t.Fatalf("Failed : %v", err)
//	}
//	resp := httptest.NewRecorder()
//	e := echo.New()
//	e.NewContext(req, resp)
//	assert.Equal(t, 501, resp.Result().StatusCode, "result is as expected")
//}
func TestController_UpdateEmployee(t *testing.T) {

}
func TestController_UpdateEmployeeFail(t *testing.T) {

}
func TestController_DeleteEmployee(t *testing.T) {
}
func TestController_DeleteEmployeeFail(t *testing.T) {

}
