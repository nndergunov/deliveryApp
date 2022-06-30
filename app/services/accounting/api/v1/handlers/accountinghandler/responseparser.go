package accountinghandler

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/accountingapi"

	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"
)

func accountToResponse(account domain.Account) accountingapi.AccountResponse {
	return accountingapi.AccountResponse{
		ID:        account.ID,
		UserID:    account.UserID,
		UserType:  account.UserType,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

func accountListToResponse(accountList []domain.Account) accountingapi.AccountListResponse {
	accountResponseList := make([]accountingapi.AccountResponse, 0, len(accountList))

	for _, account := range accountList {
		accountResponse := accountingapi.AccountResponse{
			ID:        account.ID,
			UserID:    account.UserID,
			UserType:  account.UserType,
			Balance:   account.Balance,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}

		accountResponseList = append(accountResponseList, accountResponse)
	}
	return accountingapi.AccountListResponse{
		AccountList: accountResponseList,
	}
}

func transactionToResponse(tr *domain.Transaction) accountingapi.TransactionResponse {
	return accountingapi.TransactionResponse{
		ID:            tr.ID,
		FromAccountID: tr.FromAccountID,
		ToAccountID:   tr.ToAccountID,
		Amount:        tr.Amount,
		CreatedAt:     tr.CreatedAt,
		UpdatedAt:     tr.UpdatedAt,
		Valid:         tr.Valid,
	}
}
