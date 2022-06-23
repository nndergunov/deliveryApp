package accountingclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/accountingapi"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/clients/accountingclient"
)

func TestCreateTransaction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		accountingData accountingapi.TransactionResponse
	}{
		{
			name: "creating transaction",
			accountingData: accountingapi.TransactionResponse{
				ID:            0,
				FromAccountID: 0,
				ToAccountID:   0,
				Amount:        0,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				Valid:         true,
			},
		},
	}
	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mux := http.NewServeMux()

			mux.HandleFunc("/v1/transactions", func(writer http.ResponseWriter, request *http.Request) {
				data, err := v1.Encode(test.accountingData)
				if err != nil {
					t.Fatal(err)
				}

				_, err = writer.Write(data)
				if err != nil {
					t.Fatal(err)
				}
			})

			srv := httptest.NewServer(mux)

			accntngClient := accountingclient.NewAccountingClient(srv.URL)

			valid, err := accntngClient.CreateTransaction(
				test.accountingData.ToAccountID,
				test.accountingData.FromAccountID, test.accountingData.Amount)
			if err != nil {
				t.Fatal(err)
			}

			if !valid {
				t.Fatalf("order validity: expected: %t, got: %t", test.accountingData.Valid, valid)
			}
		})
	}
}
