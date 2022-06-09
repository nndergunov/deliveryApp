package accountingclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/accountingapi"
)

type AccountingClient struct {
	accountingURL string
}

func NewAccountingClient(url string) *AccountingClient {
	return &AccountingClient{accountingURL: url}
}

func (a AccountingClient) CreateTransaction(accountID, restaurantID int, orderPrice float64) (bool, error) {
	transactionDetails, err := v1.Encode(accountingapi.TransactionRequest{
		FromAccountID: accountID,
		ToReceiverID:  restaurantID,
		ReceiverType:  "restaurant",
		Amount:        orderPrice,
	}) // Not final form, is likely to change along the course of accounting service development.
	if err != nil {
		return false, fmt.Errorf("encoding request: %w", err)
	}

	resp, err := http.Post(a.accountingURL+"/v1/transactions", "application/json", bytes.NewBuffer(transactionDetails))
	if err != nil {
		return false, fmt.Errorf("sending request: %w", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("getting response body: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return false, fmt.Errorf("closing response body: %w", err)
	}

	paymentStatus := new(accountingapi.TransactionResponse)

	err = v1.Decode(respBody, paymentStatus)
	if err != nil {
		return false, fmt.Errorf("decoding response: %w", err)
	}

	return paymentStatus.Valid, nil
}
