package couriergrp

import (
	courierapi "accounting/api/v1/accountingapi/courierrapi"
	"accounting/pkg/domain"
)

func courierAccountToResponse(courierAccount domain.CourierAccount) courierapi.CourierAccountResponse {
	return courierapi.CourierAccountResponse{
		CourierID: courierAccount.CourierID,
		Balance:    courierAccount.Balance,
	}
}
