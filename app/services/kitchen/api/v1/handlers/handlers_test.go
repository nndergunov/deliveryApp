package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/kitchenapi"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/api/v1/handlers"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"
)

var mockTasks = domain.Tasks{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}

type mockService struct{}

func (m mockService) GetTasks(_ int) (domain.Tasks, error) {
	return mockTasks, nil
}

func TestReturnTasksEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "Return tasks simple test",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mockService := new(mockService)
			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(mockService, log)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/tasks/0", nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := new(kitchenapi.Tasks)

			err := v1.Decode(resp.Body.Bytes(), respData)
			if err != nil {
				t.Fatal(err)
			}

			if len(respData.Tasks) != 5 {
				t.Errorf("Number of tasks expected: 5, got: %d", len(respData.Tasks))
			}

			for itemID, quantity := range respData.Tasks {
				if itemID < 0 || itemID > 5 {
					t.Errorf("ID out of expected range: %d", itemID)
				}

				if itemID != quantity {
					t.Errorf("itemID %d != quantity %d", itemID, quantity)
				}
			}
		})
	}
}
