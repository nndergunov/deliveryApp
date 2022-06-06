// Package restaurantgrp contains a small handlers framework extension.
package restaurantgrp

import (
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"net/http"

	"accounting/api/v1/accountingapi"
	"accounting/api/v1/accountingapi/restaurantapi"
	"accounting/pkg/service/restaurantrservice"
)

type Params struct {
	Logger            *logger.Logger
	RestaurantService restaurantservice.RestaurantService
}

// RestaurantHandler is the entrypoint into our application
type RestaurantHandler struct {
	log               *logger.Logger
	restaurantService restaurantservice.RestaurantService
}

// NewRestaurantHandler returns new http multiplexer with configured endpoints.
func NewRestaurantHandler(p Params) RestaurantHandler {

	handler := RestaurantHandler{
		log:               p.Logger,
		restaurantService: p.RestaurantService,
	}

	return handler
}

const RestaurantIDKey = "restaurant_id"

func (c RestaurantHandler) InsertNewRestaurantAccount(rw http.ResponseWriter, r *http.Request) {
	var newRestaurantAccountRequest Restaurantapi.NewRestaurantAccountRequest

	if err := accountingapi.BindJson(r, &newRestaurantAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToNewRestaurantAccount(&newRestaurantAccountRequest)

	data, err := c.restaurantService.InsertNewRestaurantAccount(account)

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
	response := RestaurantAccountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c RestaurantHandler) GetRestaurantAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[RestaurantIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoRestaurantIDParam); err != nil {
			c.log.Println(err)
		}
		return
	}

	data, err := c.restaurantService.GetRestaurantAccount(id)

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

	response := RestaurantAccountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c RestaurantHandler) DeleteRestaurantAccount(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars[RestaurantIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoRestaurantIDParam.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	data, err := c.restaurantService.DeleteRestaurantAccount(id)

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

func (c RestaurantHandler) AddBalanceRestaurantAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[RestaurantIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoRestaurantIDParam); err != nil {
			c.log.Println(err)
		}
		return
	}

	var addBalanceRestaurantAccountRequest Restaurantapi.AddRestaurantAccountRequest

	if err := accountingapi.BindJson(r, &addBalanceRestaurantAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToAddBalanceRestaurantAccount(&addBalanceRestaurantAccountRequest)

	data, err := c.restaurantService.AddToBalanceRestaurantAccount(id, account)

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

func (c RestaurantHandler) SubFromBalanceRestaurantAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[RestaurantIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoRestaurantIDParam); err != nil {
			c.log.Println(err)
		}
		return
	}

	var subBalanceRestaurantAccountRequest Restaurantapi.SubBalanceRestaurantAccountRequest

	if err := accountingapi.BindJson(r, &subBalanceRestaurantAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToSubBalanceRestaurantAccount(&subBalanceRestaurantAccountRequest)

	data, err := c.restaurantService.SubFromBalanceRestaurantAccount(id, account)

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
