package service

import (
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/pkg/repository"
)

type Model interface {
	CreateModelMain(model broker.Model) error
	CreateModelDeliveries(model broker.Model) error
	GetModel(orderUid string) (broker.Model, error)
}

type Service struct {
	Model
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Model: NewOrdersService(repos.Model),
	}
}
