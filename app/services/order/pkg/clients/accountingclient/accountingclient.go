package accountingclient

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

type AccountingClient struct {
	accountingURL string
}

func NewAccountingClient(url string) *AccountingClient {
	return &AccountingClient{accountingURL: url}
}

func (a AccountingClient) CheckIfEnoughBalance(order domain.Order) (bool, error) {
	_, err := http.Get(a.accountingURL + "v1/consumers/" + strconv.Itoa(order.FromUserID) + "/account")
	if err != nil {
		return false, fmt.Errorf("sending request: %w", err)
	}

	// TODO logic when accounting service will be in finished stage.

	return true, nil
}
