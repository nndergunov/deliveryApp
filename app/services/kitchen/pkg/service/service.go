package service

type AppService interface{}

// Service is a main service logic.
type Service struct{}

func NewService() *Service {
	return &Service{}
}
