// Package handler contains a small handler framework extension.
package consumeraccountinghandler

import (
	"accounting/api/v1/accountingapi"
	"accounting/api/v1/accountingapi/consumeraccountingapi"
	"accounting/pkg/service"
	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"io"
	"net/http"
)

type Params struct {
	Logger                    *logger.Logger
	ConsumerAccountingService service.AccountingService
}

// consumerAccountingHandler is the entrypoint into our application
type consumerAccountingHandler struct {
	serveMux          *mux.Router
	log               *logger.Logger
	accountingService service.AccountingService
}

// NewConsumerAccountingHandler returns new http multiplexer with configured endpoints.
func NewConsumerAccountingHandler(p Params) *mux.Router {
	serveMux := mux.NewRouter()

	handlerItem := consumerAccountingHandler{
		serveMux:          serveMux,
		log:               p.Logger,
		accountingService: p.ConsumerAccountingService,
	}

	handlerItem.handlerInit()

	return handlerItem.serveMux
}

const consumerIDKey = "consumer_id"

// NewConsumerAccountingHandler creates an consumerAccountingHandler value that handle a set of routes for the application.
func (c consumerAccountingHandler) handlerInit() {

	c.serveMux.HandleFunc("/status", c.StatusHandler).Methods(http.MethodPost)

	c.serveMux.HandleFunc("/v1/account/consumer", c.InsertNewConsumerAccount).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/account/consumer/{"+consumerIDKey+"}", c.GetConsumerAccount).Methods(http.MethodGet)
	c.serveMux.HandleFunc("/v1/account/consumer/{"+consumerIDKey+"}", c.DeleteConsumerAccount).Methods(http.MethodDelete)

	c.serveMux.HandleFunc("/v1/account/consumer/add", c.AddToConsumerAccount).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/account/consumer/sub", c.SubFromConsumerAccount).Methods(http.MethodPut)
}

func (c consumerAccountingHandler) StatusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
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

func (c consumerAccountingHandler) InsertNewConsumerAccount(rw http.ResponseWriter, r *http.Request) {
	var newConsumerAccountRequest consumeraccountingapi.NewConsumerAccountRequest

	if err := accountingapi.BindJson(r, &newConsumerAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToNewConsumerAccount(&newConsumerAccountRequest)

	data, err := c.accountingService.InsertNewConsumerAccount(account)

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

func (c consumerAccountingHandler) GetConsumerAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[consumerIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.accountingService.GetConsumerAccount(id)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err); err != nil {
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

func (c consumerAccountingHandler) DeleteConsumerAccount(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars[consumerIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.accountingService.DeleteConsumerAccount(id)

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

func (c consumerAccountingHandler) AddToConsumerAccount(rw http.ResponseWriter, r *http.Request) {
	var addConsumerAccountRequest consumeraccountingapi.AddConsumerAccountRequest

	if err := accountingapi.BindJson(r, &addConsumerAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToAddConsumerAccount(&addConsumerAccountRequest)

	data, err := c.accountingService.AddToConsumerAccount(account)

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

func (c consumerAccountingHandler) SubFromConsumerAccount(rw http.ResponseWriter, r *http.Request) {
	var subConsumerAccountRequest consumeraccountingapi.SubConsumerAccountRequest

	if err := accountingapi.BindJson(r, &subConsumerAccountRequest); err != nil {
		c.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	account := requestToSubConsumerAccount(&subConsumerAccountRequest)

	data, err := c.accountingService.SubFromConsumerAccount(account)

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
