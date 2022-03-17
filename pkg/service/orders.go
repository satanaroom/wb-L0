package service

import (
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/pkg/repository"
)

type OrdersService struct {
	repo  repository.Model
	cache repository.Cache
}

// Описание всех методов и функций для работы с заказами

func NewOrdersServicePostgres(repo repository.Model) *OrdersService {
	return &OrdersService{repo: repo}
}

func NewOrdersServiceCache(repo repository.Cache) *OrdersService {
	return &OrdersService{cache: repo}
}

func (s *OrdersService) CreateModelMain(model broker.Model) error {
	return s.repo.CreateModelMain(model)
}

func (s *OrdersService) CreateModelDeliveries(model broker.Model) error {
	return s.repo.CreateModelDeliveries(model)
}

func (s *OrdersService) CreateModelPayments(model broker.Model) error {
	return s.repo.CreateModelPayments(model)
}

func (s *OrdersService) CreateModelItems(model broker.Model, i int) error {
	return s.repo.CreateModelItems(model, i)
}

func (s *OrdersService) GetModel(orderUid string) (broker.Model, error) {
	return s.repo.GetModel(orderUid)
}

func (s *OrdersService) CreateModelCache(model broker.Model) error {
	return s.cache.CreateModelCache(model)
}

func (s *OrdersService) GetModelCache(orderUid string) (broker.Model, error) {
	return s.cache.GetModelCache(orderUid)
}
