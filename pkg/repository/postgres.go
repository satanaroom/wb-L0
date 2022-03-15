package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	broker "github.com/satanaroom/L0"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type BrokerPostgres struct {
	db *sqlx.DB
}

func NewPostgresDB(cfg Config) (*BrokerPostgres, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &BrokerPostgres{db: db}, nil
}

func (r *BrokerPostgres) CreateModel(model broker.Model) error {
	if model.CustomerId == "" {
		return errors.New("")
	}
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *BrokerPostgres) GetModel(orderUid string) (broker.Model, error) {
	var m broker.Model
	return m, nil
}

func (r *BrokerPostgres) CloseDB() error {
	return nil
}
