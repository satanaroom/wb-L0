package service

import (
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/pkg/repository"
)

type OrdersService struct {
	repo repository.Model
}

func NewOrdersService(repo repository.Model) *OrdersService {
	return &OrdersService{repo: repo}
}

func (s *OrdersService) CreateModelMain(model broker.Model) error {
	return s.repo.CreateModelMain(model)
}

func (s *OrdersService) CreateModelDeliveries(model broker.Model) error {
	return s.repo.CreateModelDeliveries(model)
}

func (s *OrdersService) GetModel(orderUid string) (broker.Model, error) {
	return s.repo.GetModel(orderUid)
}
