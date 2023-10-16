package test_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	restapi "github.com/dhikaroofi/simple-rest-api/internal/presentation/restApi"
	"github.com/dhikaroofi/simple-rest-api/internal/usecase"
	employee2 "github.com/dhikaroofi/simple-rest-api/internal/usecase/employee"
	"github.com/dhikaroofi/simple-rest-api/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoint_CreateEmployee(t *testing.T) {
	t.Run("Create Employee - Success", func(t *testing.T) {
		employee := employee2.Employee{
			ID:        2,
			FirstName: "bambang",
			LastName:  "purnomo",
			Email:     "purnomo@gmail.com",
			HireDate:  "2024-01-01",
		}

		employeeUseCase := &mocks.EmployeeServicesInterfaces{}
		employeeUseCase.On("Create", mock.Anything).Return(employee, nil)
		app := newServerMocks(usecase.Container{Employee: employeeUseCase})
		resp, err := sendRequest(app, "POST", "/api/v1/employee/create", employee)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		respBody, _ := decodeResponse(resp.Body)
		if respBody["status"] != http.StatusOK {
			t.Errorf("its suppose to be 200")
		}
	})

	t.Run("Create Employee - Bad Request caused by email does not match with email format", func(t *testing.T) {
		employee := employee2.Employee{
			ID:        2,
			FirstName: "bambang",
			LastName:  "purnomo",
			Email:     "purnomogmailcom",
			HireDate:  "2024-01-01",
		}

		employeeUseCase := &mocks.EmployeeServicesInterfaces{}
		employeeUseCase.On("Create", mock.Anything).Return(employee, nil)
		app := newServerMocks(usecase.Container{Employee: employeeUseCase})
		resp, err := sendRequest(app, "POST", "/api/v1/employee/create", employee)
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func decodeResponse(respBody io.ReadCloser) (resp map[string]interface{}, err error) {
	body, err := io.ReadAll(respBody)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	if err = json.Unmarshal(body, &resp); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	return
}
func newServerMocks(useCaseCont usecase.Container) *fiber.App {
	server := restapi.NewFiberServer(":8080", &useCaseCont)
	server.Route()
	return server.CallFiberApp()
}

func sendRequest(app *fiber.App, method, path string, request any) (*http.Response, error) {
	var reqBody *bytes.Reader = nil

	if request != nil {
		jsonRequestBody, _ := json.Marshal(request)
		reqBody = bytes.NewReader(jsonRequestBody)
	}

	req := httptest.NewRequest(method, path, reqBody)
	return app.Test(req, 1)
}
