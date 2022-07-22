package accountinghandler

import (
	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/rest/accountingapi"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"
)

func requestToNewAccount(req *accountingapi.NewAccountRequest) domain.Account {
	return domain.Account{
		UserID:   req.UserID,
		UserType: req.UserType,
	}
}

func requestToTransaction(req *accountingapi.TransactionRequest) domain.Transaction {
	return domain.Transaction{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
}
