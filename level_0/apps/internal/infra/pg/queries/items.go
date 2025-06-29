package queries

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/rashid567/learning-golang/level_0/apps/internal/models"
)

func SaveItems(tx *sqlx.Tx, order *models.Order) error {

	rows := make([]map[string]interface{}, 0, len(order.Items))

	for _, item := range order.Items {
		rows = append(
			rows,
			map[string]interface{}{
				"chrt_id":      item.ChrtID,
				"track_number": item.TrackNumber,
				"price":        item.Price,
				"rid":          item.Rid,
				"name":         item.Name,
				"sale":         item.Sale,
				"size":         item.Size,
				"total_price":  item.TotalPrice,
				"nm_id":        item.NmID,
				"brand":        item.Brand,
				"status":       item.Status,
				"order_uid":    order.OrderUID,
			},
		)
	}

	_, err := tx.NamedExec(
		`INSERT INTO "item" (
			chrt_id,
			track_number,
			price,
			rid,
			name,
			sale,
			size,
			total_price,
			nm_id,
			brand,
			status,
			order_uid
		) VALUES (
			:chrt_id, 
			:track_number, 
			:price, 
			:rid, 
			:name, 
			:sale, 
			:size, 
			:total_price, 
			:nm_id, 
			:brand, 
			:status, 
			:order_uid 
		) ON CONFLICT (chrt_id) DO NOTHING`,
		rows,
	)
	return err
}

func AttachItems(db *sqlx.DB, order *models.Order) error {
	err := db.Select(
		&order.Items,
		`SELECT 
			chrt_id,
			track_number,
			price,
			rid,
			name,
			sale,
			size,
			total_price,
			nm_id,
			brand,
			status
		FROM "item" 
		WHERE order_uid = $1
		ORDER BY chrt_id`,
		order.OrderUID,
	)
	if err != nil {
		return err
	}
	return nil
}
