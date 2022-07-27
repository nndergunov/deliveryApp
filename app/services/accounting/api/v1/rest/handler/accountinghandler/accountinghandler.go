package accountinghandler

import (
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	accountingapi "github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/rest/accountingapi"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/service/accountingservice"
)

type Params struct {
	Logger  *logger.Logger
	Service accountingservice.AccountService
}

// handler is the entrypoint into our application
type handler struct {
	serveMux *mux.Router
	log      *logger.Logger
	service  accountingservice.AccountService
}

// NewHandler returns new http multiplexer with configured endpoints.
func NewHandler(p Params) *mux.Router {
	serveMux := mux.NewRouter()

	handler := handler{
		serveMux: serveMux,
		log:      p.Logger,
		service:  p.Service,
	}

	handler.handlerInit()

	return handler.serveMux
}

// NewHandler creates a handler value that handle a set of routes for the application.
func (a *handler) handlerInit() {
	version := "/v1"

	a.serveMux.HandleFunc(version+"/status", a.StatusHandler).Methods(http.MethodPost)

	a.serveMux.HandleFunc(version+"/accounts", a.InsertNewAccount).Methods(http.MethodPost)
	a.serveMux.HandleFunc(version+"/accounts", a.GetAccountList).Methods(http.MethodGet)
	a.serveMux.HandleFunc(version+"/accounts/{"+accountIDKey+"}", a.GetAccount).Methods(http.MethodGet)
	a.serveMux.HandleFunc(version+"/accounts/{"+accountIDKey+"}", a.DeleteAccount).Methods(http.MethodDelete)

	a.serveMux.HandleFunc(version+"/transactions", a.InsertTransaction).Methods(http.MethodPost)

	a.serveMux.HandleFunc(version+"/transactions/{"+trIDKey+"}", a.DeleteTransaction).Methods(http.MethodDelete)
}

const (
	accountIDKey = "account_id"
	trIDKey      = "tr_id"
)

func (a handler) StatusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "accountingstorage",

		IsUp: "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		a.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		a.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	a.log.Printf("gave status %s", data.IsUp)
}

// swagger:operation POST /accounts InsertNewAccount
//
// Returns created account
//
// ---
// produces:
// - application/json
// parameters:
// - name: Body
//   in: body
//   description: account data
//   schema:
//     $ref: "#/definitions/NewAccountRequest"
//   required: true
// responses:
//   '200':
//     description: created account
//     schema:
//       $ref: "#/definitions/AccountResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (a handler) InsertNewAccount(rw http.ResponseWriter, r *http.Request) {
	var newAccountRequest accountingapi.NewAccountRequest

	if err := accountingapi.BindJson(r, &newAccountRequest); err != nil {
		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	account := requestToNewAccount(&newAccountRequest)

	data, err := a.service.InsertNewAccount(account)
	if err != nil {
		if errors.Is(err, systemErr) {
			if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				a.log.Println(err)
			}
			return
		}

		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return

	}

	response := accountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		a.log.Println(err)
		return
	}
}

// swagger:operation GET /accounts/{account_id} GetAccount
//
// Returns account
//
// ---
// produces:
// - application/json
// parameters:
// - name: account_id
//   in: path
//   description: account_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: created account
//     schema:
//       $ref: "#/definitions/AccountResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (a handler) GetAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[accountIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoIDParam); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.service.GetAccountByID(id)
	if err != nil {
		a.log.Println(err)
		if errors.Is(err, systemErr) {
			if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				a.log.Println(err)
			}
			return
		}

		if errors.Is(err, errAccountNotFound) {
			if err := accountingapi.Respond(rw, http.StatusOK, err.Error()); err != nil {
				a.log.Println(err)
			}
			return
		}

		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	response := accountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		a.log.Println(err)
		return
	}
}

// swagger:operation DELETE /accounts/{account_id} DeleteAccount
//
// Returns account
//
// ---
// produces:
// - application/json
// parameters:
// - name: account_id
//   in: path
//   description: account_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: account deleted
//     type: string
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (a handler) DeleteAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[accountIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoIDParam.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.service.DeleteAccount(id)
	if err != nil {
		a.log.Println(err)
		if errors.Is(err, systemErr) {
			if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				a.log.Println(err)
			}
			return
		}

		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return

	}

	if err := accountingapi.Respond(rw, http.StatusOK, data); err != nil {
		a.log.Println(err)
		return
	}
}

// swagger:operation POST /transactions InsertTransaction
//
// Returns account
//
// ---
// produces:
// - application/json
// parameters:
// - name: Body
//   in: body
//   description: transaction data
//   schema:
//     $ref: "#/definitions/NewAccountRequest"
//   required: true
// responses:
//   '200':
//     description: successfully transaction
//     schema:
//       $ref: "#/definitions/TransactionResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (a handler) InsertTransaction(rw http.ResponseWriter, r *http.Request) {
	var transactionRequest accountingapi.TransactionRequest

	if err := accountingapi.BindJson(r, &transactionRequest); err != nil {
		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	transaction := requestToTransaction(&transactionRequest)

	data, err := a.service.InsertTransaction(transaction)
	if err != nil {
		a.log.Println(err)

		if errors.Is(err, systemErr) {
			if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				a.log.Println(err)
			}
			return
		}

		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return

	}

	response := transactionToResponse(data)
	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		a.log.Println(err)
		return
	}
}

// swagger:operation GET /transactions GetAccountList
//
// Returns account
//
// ---
// produces:
// - application/json
// parameters:
// - type: string
//   description: user_id
//   name: user_id
//   in: query
//   required: true
// - type: string
//   description: user_id
//   name: user_id
//   in: query
//   required: true
// responses:
//   '200':
//     description:   account list data
//     schema:
//       $ref: "#/definitions/AccountListResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (a handler) GetAccountList(rw http.ResponseWriter, r *http.Request) {
	searchParam := domain.SearchParam{}

	queryParams := r.URL.Query()
	userIDList := queryParams["user_id"]
	userTypeList := queryParams["user_type"]

	if userIDList != nil && userTypeList != nil {
		userID := userIDList[0]
		userType := userTypeList[0]
		if userID != "" && userType != "" {
			searchParam["user_type"] = userType
			searchParam["user_id"] = userID
		}
	}

	data, err := a.service.GetAccountListByParam(searchParam)
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				a.log.Println(err)
			}
			return
		}

		if errors.Is(err, errAccountNotFound) {
			if err := accountingapi.Respond(rw, http.StatusOK, err.Error()); err != nil {
				a.log.Println(err)
			}
			return
		}

		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	response := accountListToResponse(data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		a.log.Println(err)
		return
	}
}

// swagger:operation DELETE /transactions/{transaction_id} DeleteTransaction
//
// Returns "transaction deleted"
//
// ---
// produces:
// - application/json
// parameters:
// - name: account_id
//   in: path
//   description: account_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: "transaction deleted"
//     schema:
//       type: string
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (a handler) DeleteTransaction(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[trIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoIDParam.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.service.DeleteTransaction(id)
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				a.log.Println(err)
			}
			return
		}

		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return

	}

	if err := accountingapi.Respond(rw, http.StatusOK, data); err != nil {
		a.log.Println(err)
		return
	}
}
