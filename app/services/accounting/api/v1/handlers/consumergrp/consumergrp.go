// Package handlers contains a small handlers framework extension.
package consumergrp

import (
	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"io"
	"net/http"

	"accounting/api/v1/accountingapi"
	"accounting/api/v1/accountingapi/consumerapi"
	"accounting/pkg/service/consumerservice"
)

type Params struct {
	Logger          *logger.Logger
	ConsumerService consumerservice.ConsumerService
}

// ConsumerHandler is the entrypoint into our application
type ConsumerHandler struct {
	log     *logger.Logger
	service consumerservice.ConsumerService
}

// NewConsumerHandler returns new http multiplexer with configured endpoints.
func NewConsumerHandler(p Params) ConsumerHandler {

	handler := ConsumerHandler{
		log:     p.Logger,
		service: p.ConsumerService,
	}

	return handler
}

const ConsumerIDKey = "consumer_id"

func (c ConsumerHandler) StatusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "accounting",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		c.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		c.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	c.log.Printf("gave status %s", data.IsUp)
}

func (c ConsumerHandler) InsertNewConsumerAccount(rw http.ResponseWriter, r *http.Request) {
	var newConsumerAccountRequest consumerapi.NewConsumerAccountRequest

	if err := accountingapi.BindJson(r, &newConsumerAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToNewConsumerAccount(&newConsumerAccountRequest)

	data, err := c.service.InsertNewConsumerAccount(account)

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
	response := consumerAccountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c ConsumerHandler) GetConsumerAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[ConsumerIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam); err != nil {
			c.log.Println(err)
		}
		return
	}

	data, err := c.service.GetConsumerAccount(id)

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

	response := consumerAccountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c ConsumerHandler) DeleteConsumerAccount(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars[ConsumerIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	data, err := c.service.DeleteConsumerAccount(id)

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

func (c ConsumerHandler) AddToBalanceConsumerAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[ConsumerIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	var addBalanceConsumerAccountRequest consumerapi.AddBalanceConsumerAccountRequest

	if err := accountingapi.BindJson(r, &addBalanceConsumerAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToAddBalanceConsumerAccount(&addBalanceConsumerAccountRequest)

	data, err := c.service.AddToBalanceConsumerAccount(id, account)

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

func (c ConsumerHandler) SubFromBalanceConsumerAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[ConsumerIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}
	var subBalanceConsumerAccountRequest consumerapi.SubBalanceConsumerAccountRequest

	if err := accountingapi.BindJson(r, &subBalanceConsumerAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToSubBalanceConsumerAccount(&subBalanceConsumerAccountRequest)

	data, err := c.service.SubFromBalanceConsumerAccount(id, account)

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
