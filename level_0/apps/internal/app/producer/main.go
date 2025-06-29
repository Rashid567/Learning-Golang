package producer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rashid567/learning-golang/level_0/apps/internal/configs"
	"github.com/rashid567/learning-golang/level_0/apps/internal/models"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/lithammer/shortuuid/v4"
	"github.com/segmentio/kafka-go"
)

type itemsStructure struct {
	Items      []models.Item
	Amount     int
	GoodsTotal int
}

func genItems() *itemsStructure {
	amount, goods_total := 0, 0
	products_count := gofakeit.Number(1, 4)
	items := make([]models.Item, 0, products_count)

	for i := 0; i < products_count; i++ {
		product := gofakeit.Product()
		price := int(product.Price)
		sale := gofakeit.Number(1, 4)
		total_price := price * sale

		amount += total_price
		goods_total += sale

		items = append(items, models.Item{
			ChrtID:      gofakeit.Int(),
			TrackNumber: shortuuid.New(),
			Price:       price,
			Rid:         shortuuid.New(),
			Name:        product.Name,
			Sale:        sale,
			Size:        "0",
			TotalPrice:  total_price,
			NmID:        gofakeit.Int(),
			Brand:       product.Benefit,
			Status:      gofakeit.Number(1, 300),
		})
	}
	return &itemsStructure{
		Amount:     amount,
		GoodsTotal: goods_total,
		Items:      items,
	}
}

func genMsg() *kafka.Message {
	order_uid := models.OrderUID(shortuuid.New())
	address := gofakeit.Address()
	items := genItems()

	value, err := json.Marshal(
		models.Order{
			OrderUID:    order_uid,
			TrackNumber: shortuuid.New(),
			Entry:       "WBIL",
			Delivery: models.Delivery{
				Name:    gofakeit.Name(),
				Phone:   gofakeit.Phone(),
				Zip:     address.Zip,
				City:    address.City,
				Address: address.Address,
				Region:  address.Country,
				Email:   gofakeit.Email(),
			},
			Payment: models.Payment{
				Transaction:  shortuuid.New(),
				RequestID:    shortuuid.New(),
				Currency:     gofakeit.RandomString([]string{"RUB", "USD"}),
				Provider:     "wbpay",
				Amount:       items.Amount,
				PaymentDt:    time.Now().Unix(),
				Bank:         gofakeit.BankName(),
				DeliveryCost: gofakeit.IntRange(100, 10000),
				GoodsTotal:   items.GoodsTotal,
				CustomFee:    0,
			},
			Items:             items.Items,
			Locale:            gofakeit.LanguageAbbreviation(),
			InternalSignature: shortuuid.New(),
			CustomerID:        shortuuid.New(),
			DeliveryService:   gofakeit.RandomString([]string{"meest", "wb"}),
			Shardkey:          "9",
			SmID:              99,
			DateCreated:       time.Now(),
			OofShard:          "1",
		},
	)
	if err != nil {
		log.Fatal("Failed created msg:", err)
	}

	return &kafka.Message{
		Key:   []byte(order_uid),
		Value: value,
	}

}

func sendMsg(ctx context.Context, writer *kafka.Writer, msg *kafka.Message) {
	err := writer.WriteMessages(ctx, *msg)
	if err != nil {
		log.Fatal("Failed to send msgs:", err)
	}
	log.Printf("Succesfully sent msg %s\n", msg.Key)
}

func Run() {
	conf := configs.GetConfig[configs.ProducerConfig]()
	ctx := context.Background()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: conf.KafkaProducer.Hosts(),
		Topic:   conf.KafkaProducer.Topic,
	})
	defer writer.Close()

	for {
		msg := genMsg()
		sendMsg(ctx, writer, msg)
		time.Sleep(time.Second)
	}

}
