package handler

import (
	"context"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/grpc/proto"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/service/accountingservice"
)

type Params struct {
	Logger         *logger.Logger
	AccountService accountingservice.AccountService
}

// handler is the entrypoint into our application
type handler struct {
	pb.UnimplementedAccountingServer
	service accountingservice.AccountService
	log     *logger.Logger
}

// NewHandler returns new http multiplexer with configured endpoints.
func NewHandler(p Params) *grpc.Server {
	h := &handler{
		log:     p.Logger,
		service: p.AccountService,
	}

	srv := grpc.NewServer()
	pb.RegisterAccountingServer(srv, h)

	return srv
}

// InsertNewAccount implements accounting_proto.AccountingServer
func (h *handler) InsertNewAccount(ctx context.Context, in *pb.NewAccountRequest) (*pb.AccountResponse, error) {
	account := domain.Account{
		UserID:   int(in.UserID),
		UserType: in.UserType,
	}

	newAccount, err := h.service.InsertNewAccount(account)
	if err != nil {
		return nil, err
	}

	return &pb.AccountResponse{
		ID:        int64(newAccount.ID),
		UserID:    int64(newAccount.UserID),
		UserType:  newAccount.UserType,
		Balance:   newAccount.Balance,
		CreatedAt: timestamppb.New(newAccount.CreatedAt),
		UpdatedAt: timestamppb.New(newAccount.UpdatedAt),
	}, nil
}

// GetAccountList implements accounting_proto.AccountingServer
func (h *handler) GetAccountList(ctx context.Context, in *pb.SearchParam) (*pb.AccountListResponse, error) {
	searchParam := domain.SearchParam{}
	searchParam["user_type"] = *in.UserType
	searchParam["user_id"] = strconv.Itoa(int(*in.UserID))

	outList, err := h.service.GetAccountListByParam(searchParam)
	if err != nil {
		return nil, err
	}

	if outList == nil {
		return nil, nil
	}

	var responses []*pb.AccountResponse

	for _, out := range outList {
		resp := &pb.AccountResponse{
			ID:        int64(out.ID),
			UserID:    int64(out.UserID),
			UserType:  out.UserType,
			Balance:   out.Balance,
			CreatedAt: timestamppb.New(out.CreatedAt),
			UpdatedAt: timestamppb.New(out.UpdatedAt),
		}
		responses = append(responses, resp)
	}

	return &pb.AccountListResponse{AccountList: responses}, nil
}

// GetAccount implements accounting_proto.AccountingServer
func (h *handler) GetAccount(ctx context.Context, in *pb.AccountID) (*pb.AccountResponse, error) {
	out, err := h.service.GetAccountByID(strconv.Itoa(int(in.AccountID)))
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, nil
	}

	return &pb.AccountResponse{
		ID:        int64(out.ID),
		UserID:    int64(out.UserID),
		UserType:  out.UserType,
		Balance:   out.Balance,
		CreatedAt: timestamppb.New(out.CreatedAt),
		UpdatedAt: timestamppb.New(out.UpdatedAt),
	}, nil
}

// DeleteAccount implements accounting_proto.AccountingServer
func (h *handler) DeleteAccount(ctx context.Context, in *pb.AccountID) (*pb.AccountDeleteResponse, error) {
	out, err := h.service.DeleteAccount(strconv.Itoa(int(in.AccountID)))
	if err != nil {
		return nil, err
	}

	return &pb.AccountDeleteResponse{AccountDeleteResponse: out}, nil
}

// InsertTransaction implements accounting_proto.AccountingServer
func (h *handler) InsertTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	account := domain.Transaction{
		FromAccountID: int(in.FromAccountID),
		ToAccountID:   int(in.ToAccountID),
		Amount:        in.Amount,
	}

	out, err := h.service.InsertTransaction(account)
	if err != nil {
		return nil, err
	}

	return &pb.TransactionResponse{
		ID:            int64(out.ID),
		FromAccountID: int64(out.FromAccountID),
		ToAccountID:   int64(out.ToAccountID),
		Amount:        out.Amount,
		CreatedAt:     timestamppb.New(out.CreatedAt),
		UpdatedAt:     timestamppb.New(out.UpdatedAt),
		Valid:         out.Valid,
	}, nil
}

// DeleteTransaction implements accounting_proto.AccountingServer
func (h *handler) DeleteTransaction(ctx context.Context, in *pb.TransactionID) (*pb.TransactionDeleteResponse, error) {
	out, err := h.service.DeleteTransaction(strconv.Itoa(int(in.TransactionID)))
	if err != nil {
		return nil, err
	}

	return &pb.TransactionDeleteResponse{TransactionDeleteResponse: out}, nil
}
