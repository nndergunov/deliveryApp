// Package couriergrp contains a small handlers framework extension.
package couriergrp

import (
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"net/http"

	"accounting/api/v1/accountingapi"
	"accounting/api/v1/accountingapi/courierrapi"
	"accounting/pkg/service/courierservice"
)

type Params struct {
	Logger         *logger.Logger
	CourierService courierservice.CourierService
}

// CourierHandler is the entrypoint into our application
type CourierHandler struct {
	log     *logger.Logger
	service courierservice.CourierService
}

// NewCourierHandler returns new http multiplexer with configured endpoints.
func NewCourierHandler(p Params) CourierHandler {

	handler := CourierHandler{
		log:     p.Logger,
		service: p.CourierService,
	}

	return handler
}

const CourierIDKey = "courier_id"

func (c CourierHandler) InsertNewCourierAccount(rw http.ResponseWriter, r *http.Request) {
	var newCourierAccountRequest courierapi.NewCourierAccountRequest

	if err := accountingapi.BindJson(r, &newCourierAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToNewCourierAccount(&newCourierAccountRequest)

	data, err := c.service.InsertNewCourierAccount(account)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}
	response := courierAccountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
	}
}

func (c CourierHandler) GetCourierAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[CourierIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam); err != nil {
			c.log.Println(err)
		}
		return
	}

	data, err := c.service.GetCourierAccount(id)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierAccountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c CourierHandler) DeleteCourierAccount(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars[CourierIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	data, err := c.service.DeleteCourierAccount(id)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err := accountingapi.Respond(rw, http.StatusOK, data); err != nil {
		c.log.Println(err)
		return
	}
}

func (c CourierHandler) AddToBalanceCourierAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[CourierIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam); err != nil {
			c.log.Println(err)
		}
		return
	}

	var addBalanceCourierAccountRequest courierapi.AddBalanceCourierAccountRequest

	if err := accountingapi.BindJson(r, &addBalanceCourierAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToAddBalanceCourierAccount(&addBalanceCourierAccountRequest)

	data, err := c.service.AddToBalanceCourierAccount(id, account)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err := accountingapi.Respond(rw, http.StatusOK, data); err != nil {
		c.log.Println(err)
		return
	}
}

func (c CourierHandler) SubFromBalanceCourierAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[CourierIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam); err != nil {
			c.log.Println(err)
		}
		return
	}

	var subBalanceCourierAccountRequest courierapi.SubBalanceCourierAccountRequest

	if err := accountingapi.BindJson(r, &subBalanceCourierAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToSubBalanceCourierAccount(&subBalanceCourierAccountRequest)

	data, err := c.service.SubFromBalanceCourierAccount(id, account)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err := accountingapi.Respond(rw, http.StatusOK, data); err != nil {
		c.log.Println(err)
		return
	}
}
