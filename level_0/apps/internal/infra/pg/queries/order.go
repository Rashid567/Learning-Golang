package queries

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/rashid567/learning-golang/level_0/apps/internal/models"
	"github.com/rashid567/learning-golang/level_0/apps/internal/utils"
)

func SaveOrder(tx *sqlx.Tx, order *models.Order) error {
	ord := *order

	delivery, err := json.Marshal(ord.Delivery)
	if err != nil {
		return err
	}

	_, err = tx.NamedExec(
		`INSERT INTO "order" (
			order_uid,
			track_number,
			entry,
			delivery,
			locale,
			internal_signature,
			customer_id,
			delivery_service,
			shardkey,
			sm_id,
			date_created,
			oof_shard
		) VALUES (
			:order_uid,
			:track_number,
			:entry,
			:delivery,
			:locale,
			:internal_signature,
			:customer_id,
			:delivery_service,
			:shardkey,
			:sm_id,
			:date_created,
			:oof_shard
		) ON CONFLICT (order_uid) DO NOTHING`,
		map[string]interface{}{
			"order_uid":          ord.OrderUID,
			"track_number":       ord.TrackNumber,
			"entry":              ord.Entry,
			"delivery":           delivery,
			"locale":             ord.Locale,
			"internal_signature": ord.InternalSignature,
			"customer_id":        ord.CustomerID,
			"delivery_service":   ord.DeliveryService,
			"shardkey":           ord.Shardkey,
			"sm_id":              ord.SmID,
			"date_created":       ord.DateCreated,
			"oof_shard":          ord.OofShard,
		},
	)
	return err
}

func GetOrder(db *sqlx.DB, orderUID models.OrderUID) (*models.Order, error) {
	var order models.Order

	err := db.Get(
		&order,
		`SELECT * FROM "order" WHERE order_uid = $1`,
		orderUID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.AppError{
				Detail:     "Order not found",
				StatusCode: http.StatusNotFound,
			}
		}
		return nil, err
	}
	return &order, nil
}

func GetOrderUIDs(db *sqlx.DB, limit int) (res []models.OrderUID, err error) {
	err = db.Select(
		&res,
		`SELECT order_uid FROM "order" 
		ORDER BY date_created DESC LIMIT $1`,
		limit,
	)
	if err != nil {
		return
	}
	fmt.Printf("Got %d order ids\n", len(res))
	return
}
