package repository

import broker "github.com/satanaroom/L0"

type Broker interface {
	CreateModel(model broker.Model) error
	GetModel(orderUid string) (broker.Model, error)
	CloseDB() error
}

