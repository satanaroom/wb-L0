package repository

import (
	"time"

	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/cache"
)

type OrderCache struct {
	cache *cache.Cache
}

func NewOrderCache(cache *cache.Cache) *OrderCache {
	return &OrderCache{cache: cache}
}

// Метод записи в кэш основных данных о заказе
func (r *OrderCache) CreateModelCache(model broker.Model) error {
	r.cache.Set(model.OrderUid, model, 5*time.Minute)
	return nil
}

// Метод получения из кэша данных о заказе по его id
func (r *OrderCache) GetModelCache(orderUid string) (broker.Model, error) {
	var empty broker.Model
	val, _ := r.cache.Get(orderUid)
	switch v := val.(type) {
	case broker.Model:
		return v, nil
	case nil:
		return empty, nil
	}
	return empty, nil
}
