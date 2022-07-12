package handlers

import (
	"github.com/nndergunov/deliveryApp/app/services/kitchen/api/v1/communication"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"
)

func tasksToResponse(tasks domain.Tasks) communication.Tasks {
	return communication.Tasks{
		Tasks: tasks,
	}
}
