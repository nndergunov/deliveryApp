// Package accountingclient is responsible for communication with accounting service.
package accountingclient

import (
	"bytes"
	"fmt"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/accountingapi"
)

// AccountingClient communicates with accounting service.
type AccountingClient struct {
	accountingURL string
}

// NewAccountingClient returns new AccountingClient instance.
func NewAccountingClient(url string) *AccountingClient {
	return &AccountingClient{accountingURL: url}
}

// CreateTransaction sends transaction data to the accounting service.
func (a AccountingClient) CreateTransaction(accountID, restaurantID int, orderPrice float64) (bool, error) {
	transactionDetails, err := v1.Encode(accountingapi.TransactionRequest{
		FromAccountID: accountID,
		ToAccountID:   restaurantID,
		Amount:        orderPrice,
	}) // Not final form, is likely to change along the course of accounting service development.
	if err != nil {
		return false, fmt.Errorf("encoding request: %w", err)
	}

	resp, err := http.Post(a.accountingURL+"/v1/transactions", "application/json", bytes.NewBuffer(transactionDetails))
	if err != nil {
		return false, fmt.Errorf("sending request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("%w: error %d", ErrAccountingServiceFail, resp.StatusCode)
	}

	paymentStatus := new(accountingapi.TransactionResponse)

	err = v1.DecodeResponse(resp, paymentStatus)
	if err != nil {
		return false, fmt.Errorf("decoding response: %w", err)
	}

	return true, nil
}
