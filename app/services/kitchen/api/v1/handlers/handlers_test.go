package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/api/v1/handlers"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/api/v1/kitchenapi"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/service/mockservice"
	"github.com/stretchr/testify/mock"
)

func TestReturnTasksEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		restaurantID int
		tasks        domain.Tasks
	}{
		{
			name:         "Return tasks simple",
			restaurantID: 0,
			tasks:        domain.Tasks{1: 1, 2: 2, 3: 3, 4: 4, 5: 5},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var repo mockservice.AppService

			repo.On("GetTasks", mock.AnythingOfType("int")).Return(test.tasks, nil).Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/tasks/"+strconv.Itoa(test.restaurantID), nil)

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
