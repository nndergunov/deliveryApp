package couriergrp

import (
	"accounting/api/v1/accountingapi/courierrapi"
	"accounting/pkg/domain"
)

func requestToNewCourierAccount(req *courierapi.NewCourierAccountRequest) domain.CourierAccount {
	return domain.CourierAccount{
		CourierID: req.CourierID,
	}
}

func requestToAddBalanceCourierAccount(req *courierapi.AddBalanceCourierAccountRequest) domain.CourierAccount {
	return domain.CourierAccount{
		CourierID: req.CourierID,
		Balance:   req.Amount,
	}
}

func requestToSubBalanceCourierAccount(req *courierapi.SubBalanceCourierAccountRequest) domain.CourierAccount {
	return domain.CourierAccount{
		CourierID: req.CourierID,
		Balance:   req.Amount,
	}
}
