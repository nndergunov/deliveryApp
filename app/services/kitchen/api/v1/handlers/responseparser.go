package handlers

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/kitchenapi"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"
)

func tasksToResponse(tasks domain.Tasks) kitchenapi.Tasks {
	return kitchenapi.Tasks{
		Tasks: tasks,
	}
}
