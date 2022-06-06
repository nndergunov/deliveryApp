package consumergrp

import (
	"accounting/api/v1/accountingapi/consumerapi"
	"accounting/pkg/domain"
)

func requestToNewConsumerAccount(req *consumerapi.NewConsumerAccountRequest) domain.ConsumerAccount {
	return domain.ConsumerAccount{
		ConsumerID: req.ConsumerID,
	}
}

func requestToAddBalanceConsumerAccount(req *consumerapi.AddBalanceConsumerAccountRequest) domain.ConsumerAccount {
	return domain.ConsumerAccount{
		Balance: req.Amount,
	}
}

func requestToSubBalanceConsumerAccount(req *consumerapi.SubBalanceConsumerAccountRequest) domain.ConsumerAccount {
	return domain.ConsumerAccount{
		Balance: req.Amount,
	}
}
