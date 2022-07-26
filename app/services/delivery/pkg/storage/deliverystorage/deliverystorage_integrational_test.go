package deliverystorage_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/docker"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/storage/deliverystorage"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/db/dbtest"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"
)

var c *docker.Container

func TestMain(m *testing.M) {
	var err error
	c, err = dbtest.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtest.StopDB(c)

	m.Run()
}

func TestAssignOrder(t *testing.T) {
	tests := []struct {
		name string
		in   domain.AssignOrder
		out  *domain.AssignOrder
	}{
		{
			"test_assign_order",
			domain.AssignOrder{
				OrderID:   1,
				CourierID: 2,
			},
			&domain.AssignOrder{
				OrderID:   1,
				CourierID: 2,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := deliverystorage.NewStorage(deliverystorage.Params{DB: database})

			resp, err := s.AssignOrder(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, resp, test.out)
		})
	}
}
