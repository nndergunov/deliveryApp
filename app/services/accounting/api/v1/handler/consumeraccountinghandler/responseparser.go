package consumeraccountinghandler

import (
	"accounting/api/v1/accountingapi/consumeraccountingapi"
	"accounting/pkg/domain"
)

func consumerAccountToResponse(consumerAccount domain.ConsumerAccount) consumeraccountingapi.ConsumerAccountResponse {
	return consumeraccountingapi.ConsumerAccountResponse{
		ConsumerID: consumerAccount.ConsumerID,
		Balance:    consumerAccount.Balance,
	}
}
