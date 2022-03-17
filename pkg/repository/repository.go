package repository

import (
	"github.com/jmoiron/sqlx"
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/cache"
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

type Repository struct {
	Model
	Cache
}

// Функция инициализации репозитория
func NewRepository(db *sqlx.DB, cache *cache.Cache) *Repository {
	return &Repository{
		Model: NewOrderPostgres(db),
		Cache: NewOrderCache(cache),
	}
}
