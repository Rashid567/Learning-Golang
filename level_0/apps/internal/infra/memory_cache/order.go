package memory_cache

import (
	"log"

	"github.com/rashid567/learning-golang/level_0/apps/internal/domain/order"
	"github.com/rashid567/learning-golang/level_0/apps/internal/models"
)

func GetOrderCache() order.OrderCache {
	return &orderCache{
		cache: make(map[models.OrderUID]models.Order),
	}
}

type orderCache struct {
	cache map[models.OrderUID]models.Order
}

func (s *orderCache) Set(order *models.Order) error {
	s.cache[order.OrderUID] = *order
	log.Printf("Order %s updated in cache\n", order.OrderUID)
	return nil
}

func (s *orderCache) Get(orderUID models.OrderUID) (*models.Order, error) {
	if order, ok := s.cache[orderUID]; ok {
		log.Printf("Order %s found in cache\n", orderUID)
		return &order, nil
	}
	log.Printf("Order %s NOT found in cache\n", orderUID)
	return nil, nil
}
