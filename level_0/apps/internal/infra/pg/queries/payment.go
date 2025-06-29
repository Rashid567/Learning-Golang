package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/rashid567/learning-golang/level_0/apps/internal/models"
)

func SavePayment(tx *sqlx.Tx, order *models.Order) error {
	payment := order.Payment

	_, err := tx.NamedExec(
		`INSERT INTO payment (
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee,
			order_uid
		) VALUES (
			:transaction,
			:request_id,
			:currency,
			:provider,
			:amount,
			:payment_dt,
			:bank,
			:delivery_cost,
			:goods_total,
			:custom_fee,
			:order_uid
		) ON CONFLICT (transaction) DO NOTHING`,
		map[string]interface{}{
			"transaction":   payment.Transaction,
			"request_id":    payment.RequestID,
			"currency":      payment.Currency,
			"provider":      payment.Provider,
			"amount":        payment.Amount,
			"payment_dt":    payment.PaymentDt,
			"bank":          payment.Bank,
			"delivery_cost": payment.DeliveryCost,
			"goods_total":   payment.GoodsTotal,
			"custom_fee":    payment.CustomFee,
			"order_uid":     order.OrderUID,
		},
	)
	return err
}

func AttachPayment(db *sqlx.DB, order *models.Order) error {
	log.Println(order.OrderUID)
	err := db.Get(
		&order.Payment,
		`SELECT 
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee
		FROM "payment" 
		WHERE order_uid = $1`,
		order.OrderUID,
	)
	if err != nil {
		return err
	}
	return nil
}
