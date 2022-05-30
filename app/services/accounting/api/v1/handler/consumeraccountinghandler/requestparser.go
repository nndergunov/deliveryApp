package consumeraccountinghandler

import (
	"accounting/api/v1/accountingapi/consumeraccountingapi"
	"accounting/pkg/domain"
)

func requestToNewConsumerAccount(req *consumeraccountingapi.NewConsumerAccountRequest) domain.ConsumerAccount {
	return domain.ConsumerAccount{
		ConsumerID: req.ConsumerID,
	}
}

func requestToAddConsumerAccount(req *consumeraccountingapi.AddConsumerAccountRequest) domain.ConsumerAccount {
	return domain.ConsumerAccount{
		ConsumerID: req.ConsumerID,
		Balance:    req.Amount,
	}
}

func requestToSubConsumerAccount(req *consumeraccountingapi.SubConsumerAccountRequest) domain.ConsumerAccount {
	return domain.ConsumerAccount{
		ConsumerID: req.ConsumerID,
		Balance:    req.Amount,
	}
}
