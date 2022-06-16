package tokenservice

import (
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/nndergunov/delivryApp/app/services/accounting/pkg/domain"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	Keys string
}

type tokenService struct {
	key string
}

// TokenService is the interface for the auth service.
type TokenService interface {
	CreateToken(tr *domain.Transaction) (*string, error)
	getTokens(privateCl interface{}) (string, error)
}

// NewTokenService constructs a new AuthService.
func NewTokenService(p Params) (TokenService, error) {
	tokenServiceItem := &tokenService{
		key: p.Keys,
	}

	return tokenServiceItem, nil
}

func (t *tokenService) CreateToken(tr *domain.Transaction) (*string, error) {
	privateCl := struct {
		Transaction *domain.Transaction `json:"transaction"`
	}{
		tr,
	}

	token, err := t.getTokens(privateCl)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (t *tokenService) getTokens(privateCl interface{}) (string, error) {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS512, Key: t.key},
		(&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := jwt.Claims{
		Issuer:   "accountingservice",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Subject:  "transaction token",
	}

	token, err := jwt.Signed(sig).Claims(cl).Claims(privateCl).CompactSerialize()
	if err != nil {
		return "", err
	}
	return token, nil
}
