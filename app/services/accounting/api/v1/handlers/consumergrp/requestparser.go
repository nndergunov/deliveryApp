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

func requestToAddConsumerAccount(req *consumerapi.AddConsumerAccountRequest) domain.ConsumerAccount {
	return domain.ConsumerAccount{
		ConsumerID: req.ConsumerID,
		Balance:    req.Amount,
	}
}

func requestToSubConsumerAccount(req *consumerapi.SubConsumerAccountRequest) domain.ConsumerAccount {
	return domain.ConsumerAccount{
		ConsumerID: req.ConsumerID,
		Balance:    req.Amount,
	}
}
