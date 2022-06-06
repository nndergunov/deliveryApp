package consumergrp

import (
	"accounting/api/v1/accountingapi/consumerapi"
	"accounting/pkg/domain"
)

func consumerAccountToResponse(consumerAccount domain.ConsumerAccount) consumerapi.ConsumerAccountResponse {
	return consumerapi.ConsumerAccountResponse{
		ConsumerID: consumerAccount.ConsumerID,
		Balance:    consumerAccount.Balance,
	}
}
