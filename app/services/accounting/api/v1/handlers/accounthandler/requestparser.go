package accounthandler

import (
	"accounting/api/v1/accountingapi"
	"accounting/pkg/domain"
)

func requestToNewAccount(req *accountingapi.UserRequest) domain.Account {
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
