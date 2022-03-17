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

// Метод записи в БД таблицы с основными данными о заказе
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
		logrus.Printf("[db] orders %s", err)
		tx.Rollback()
	}
	return nil
}

// Метод записи в БД таблицы с данными о доставке заказа
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
		logrus.Printf("[db] deliveries %s", err)
		tx.Rollback()
		return nil
	}
	logrus.Println("[db] table deliveries created")
	return tx.Commit()
}

// Метод записи в БД таблицы с платежными данными заказа
func (r *OrderPostgres) CreateModelPayments(model broker.Model) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Fatalf("[db] failed to tx begin: %s", err.Error())
		return err
	}

	createModelPaymentsQuery := fmt.Sprintf(`
		INSERT INTO %s
		(transaction, request_id, currency,
		provider, amount, payment_dt, bank,
		delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, paymentsTable)

	_, err = tx.Exec(createModelPaymentsQuery, model.OrderUid, model.Payment.RequestId,
		model.Payment.Currency, model.Payment.Provider, model.Payment.Amount,
		model.Payment.PaymentDt, model.Payment.Bank, model.Payment.DeliveryCost,
		model.Payment.GoodsTotal, model.Payment.CustomFee)
	if err != nil {
		logrus.Printf("[db] payments %s", err)
		tx.Rollback()
		return nil
	}
	logrus.Println("[db] table payments created")
	return tx.Commit()
}

// Метод записи в БД таблицы с данными об элементах заказа
func (r *OrderPostgres) CreateModelItems(model broker.Model, i int) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Fatalf("[db] failed to tx begin: %s", err.Error())
		return err
	}

	createModelItemsQuery := fmt.Sprintf(`
		INSERT INTO %s
		(order_uid, chrt_id, track_number, price,
		rid, name, sale, size,
		total_price, nm_id, brand, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`, itemsTable)

	_, err = tx.Exec(createModelItemsQuery, model.OrderUid, model.Items[i].ChrtId,
		model.Items[i].TrackNumber, model.Items[i].Price, model.Items[i].Rid,
		model.Items[i].Name, model.Items[i].Sale, model.Items[i].Size,
		model.Items[i].TotalPrice, model.Items[i].NmId, model.Items[i].Brand, model.Items[i].Status)
	if err != nil {
		logrus.Printf("[db] items %s", err)
		tx.Rollback()
		return nil
	}

	logrus.Println("[db] table items created")
	return tx.Commit()
}

// Метод получения из БД таблиц о заказе по его id
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
	getModelPaymentsQuery := fmt.Sprintf("SELECT * FROM %s WHERE transaction = $1", paymentsTable)
	err = r.db.Get(&model.Payment, getModelPaymentsQuery, orderUid)
	if err != nil {
		return model, err
	}
	items := model.Items
	getModelItemsQuery := fmt.Sprintf("SELECT * FROM %s WHERE order_uid = $1", itemsTable)
	err = r.db.Select(&items, getModelItemsQuery, orderUid)
	if err != nil {
		model.Items = items
		return model, err
	}
	model.Items = items
	logrus.Println("[db] selected success")
	return model, nil
}
