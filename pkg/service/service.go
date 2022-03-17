package service

import (
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/pkg/repository"
)

// Описание интерфейса для работы с БД
type Model interface {
	CreateModelMain(model broker.Model) error
	CreateModelDeliveries(model broker.Model) error
	CreateModelPayments(model broker.Model) error
	CreateModelItems(model broker.Model, i int) error
	GetModel(orderUid string) (broker.Model, error)
}

// Описание интерфейса для работы с кэшем
type Cache interface {
	CreateModelCache(model broker.Model) error
	GetModelCache(orderUid string) (broker.Model, error)
}

type Service struct {
	Model
	Cache
}

// Инициализация сервисов
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Model: NewOrdersServicePostgres(repos.Model),
		Cache: NewOrdersServiceCache(repos.Cache),
	}
}
