package repository

import (
	"github.com/jmoiron/sqlx"
	broker "github.com/satanaroom/L0"
)

type Model interface {
	CreateModelMain(model broker.Model) error
	CreateModelDeliveries(model broker.Model) error
	GetModel(orderUid string) (broker.Model, error)
}

type Repository struct {
	Model
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Model: NewOrderPostgres(db),
	}
}
