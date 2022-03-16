package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	broker "github.com/satanaroom/L0"
	"github.com/sirupsen/logrus"
)

const (
	ordersTable     = "orders"
	deliveriesTable = "deliveries"
	paymentsTable   = "payments"
	itemsTable      = "items"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateModelMain(model broker.Model) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Fatalf("[db] failed to tx begin: %s", err.Error())
		return err
	}

	createModelMainQuery := fmt.Sprintf(`
		INSERT INTO %s (order_uid, track_number, entry, locale,
		internal_signature, customer_id, delivery_service,
		shardkey, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, ordersTable)

	row := tx.QueryRow(createModelMainQuery, model.OrderUid, model.TrackNumber,
		model.Entry, model.Locale, model.InternalSignature,
		model.CustomerId, model.DeliveryService, model.Shardkey,
		model.SmId, model.DateCreated, model.OofShard)

	err = row.Scan(&model.OrderUid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			logrus.Println("[db] table orders created")
			return tx.Commit()
		}
		logrus.Println("[db] orders is published yet")
		tx.Rollback()
	}
	return nil
}

func (r *OrderPostgres) CreateModelDeliveries(model broker.Model) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Fatalf("[db] failed to tx begin: %s", err.Error())
		return err
	}

	createModelDeliveriesQuery := fmt.Sprintf(`
		INSERT INTO %s
		(order_uid, name, phone,
		zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, deliveriesTable)

	_, err = tx.Exec(createModelDeliveriesQuery, model.OrderUid, model.Delivery.Name,
		model.Delivery.Phone, model.Delivery.Zip, model.Delivery.City,
		model.Delivery.Address, model.Delivery.Region, model.Delivery.Email)
	if err != nil {
		logrus.Println("[db] deliveries is published yet")
		tx.Rollback()
		return nil
	}
	logrus.Println("[db] table deliveries created")
	return tx.Commit()
}

func (r *OrderPostgres) GetModel(orderUid string) (broker.Model, error) {
	getModelMainQuery := fmt.Sprintf("SELECT * FROM %s WHERE order_uid = $1", ordersTable)
	var model broker.Model
	err := r.db.Get(&model, getModelMainQuery, orderUid)
	if err != nil {
		return model, err
	}
	getModelDeliveriesQuery := fmt.Sprintf("SELECT * FROM %s WHERE order_uid = $1", deliveriesTable)
	err = r.db.Get(&model.Delivery, getModelDeliveriesQuery, orderUid)
	if err != nil {
		return model, err
	}
	logrus.Println("[db] selected success")
	return model, nil
}
