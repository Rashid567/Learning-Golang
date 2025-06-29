package pg

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/rashid567/learning-golang/level_0/apps/internal/configs"
	"github.com/rashid567/learning-golang/level_0/apps/internal/domain/order"
	"github.com/rashid567/learning-golang/level_0/apps/internal/infra/pg/queries"
	"github.com/rashid567/learning-golang/level_0/apps/internal/models"
)

func GetOrderRepo(conf *configs.PostgresConfig) order.OrderRepo {
	conn_string := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.Host, conf.Port, conf.Login, conf.Pasword, conf.DatabaseName, conf.SSLMode,
	)
	log.Println(conn_string)

	db, err := sqlx.Connect("postgres", conn_string)
	if err != nil {
		log.Fatalln(err)
	}

	return &orderRepo{db: db}
}

type orderRepo struct {
	db *sqlx.DB
}

func (s *orderRepo) SaveOrder(order *models.Order) error {
	tx := s.db.MustBegin()
	defer tx.Rollback()

	err := queries.SaveOrder(tx, order)
	if err != nil {
		return err
	}

	err = queries.SavePayment(tx, order)
	if err != nil {
		return err
	}

	err = queries.SaveItems(tx, order)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (s *orderRepo) GetOrder(orderUID models.OrderUID) (*models.Order, error) {

	order, err := queries.GetOrder(s.db, orderUID)
	if err != nil {
		return nil, err
	}

	err = queries.AttachPayment(s.db, order)
	if err != nil {
		return nil, err
	}

	err = queries.AttachItems(s.db, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderRepo) GetAllOrders(limit int) (*[]models.Order, error) {
	orders := make([]models.Order, 0, limit)

	orderUIDs, err := queries.GetOrderUIDs(s.db, limit)
	if err != nil {
		return nil, err
	}

	for _, orderUID := range orderUIDs {
		order, err := queries.GetOrder(s.db, orderUID)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}

	return &orders, nil
}
