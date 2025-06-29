package order

import (
	"log"

	"github.com/rashid567/learning-golang/level_0/apps/internal/models"
)

type OrderRepo interface {
	SaveOrder(order *models.Order) error
	GetOrder(orderUID models.OrderUID) (*models.Order, error)
	GetAllOrders(limit int) (*[]models.Order, error)
}

type OrderCache interface {
	Set(order *models.Order) error
	Get(orderUID models.OrderUID) (*models.Order, error)
}

type OrderService interface {
	SaveOrder(order *models.Order) error
	GetOrder(orderUID models.OrderUID) (*models.Order, error)
	RestoreCache()
}

// NewOrderService создает новый экземпляр сервиса
func NewOrderService(repo *OrderRepo, cache *OrderCache) OrderService {
	return &service{
		repo:  *repo,
		cache: *cache,
	}
}

type service struct {
	repo  OrderRepo
	cache OrderCache
}

func (s *service) SaveOrder(order *models.Order) error {
	// Сохраняем в БД в транзакции
	err := s.repo.SaveOrder(order)
	if err != nil {
		log.Printf("Failed to save order %s: %v", order.OrderUID, err)
		return err
	}

	s.cache.Set(order)

	log.Printf("Order %s saved successfully", order.OrderUID)
	return nil
}

func (s *service) GetOrder(orderUID models.OrderUID) (*models.Order, error) {
	cached_order, err := s.cache.Get(orderUID)
	if cached_order != nil {
		log.Println("Got order from cache")
		return cached_order, nil
	} else if err != nil {
		log.Printf("Failed to get order from cache: %s\n", err)
	}

	order, err := s.repo.GetOrder(orderUID)
	if err != nil {
		return nil, err
	}

	s.cache.Set(order)

	log.Println("Gor order from repo")
	return order, nil
}

func (s *service) RestoreCache() {
	orders, err := s.repo.GetAllOrders(10)
	if err != nil {
		log.Fatalf("Failed to get orders: %s", err)
	}

	for _, order := range *orders {
		s.cache.Set(&order)
	}

	log.Printf("Cache restored with %d orders", len(*orders))
}
