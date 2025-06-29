package web_and_consumer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v3"
	fiber_compress "github.com/gofiber/fiber/v3/middleware/compress"
	fiber_logger "github.com/gofiber/fiber/v3/middleware/logger"
	fiber_recover "github.com/gofiber/fiber/v3/middleware/recover"
	fiber_static "github.com/gofiber/fiber/v3/middleware/static"

	"github.com/rashid567/learning-golang/level_0/apps/internal/api/handlers"
	"github.com/rashid567/learning-golang/level_0/apps/internal/configs"
	"github.com/rashid567/learning-golang/level_0/apps/internal/domain/order"
	"github.com/rashid567/learning-golang/level_0/apps/internal/infra/memory_cache"
	"github.com/rashid567/learning-golang/level_0/apps/internal/infra/pg"
)

func config_fiber_app(app *fiber.App, orderService order.OrderService) {

	app.Use(fiber_recover.New())
	app.Use(fiber_compress.New())
	app.Use(fiber_logger.New(fiber_logger.Config{
		TimeFormat: "2006/01/02 15:04:05",
	}))
	app.Use("/", fiber_static.New("../../internal/web/static/"))

	orderHandler := handlers.NewHandler(orderService)
	orderHandler.Register(app)

}

func getOrderService(ctx context.Context, conf *configs.WebAndConsumerConfig) order.OrderService {
	order_cache := memory_cache.GetOrderCache()
	order_repo := pg.GetOrderRepo(&conf.Postgres)
	res := order.NewOrderService(&order_repo, &order_cache)
	go res.RestoreCache()
	return res
}

func Run() {
	var wg sync.WaitGroup
	conf := configs.GetConfig[configs.WebAndConsumerConfig]()
	ctx, cancel := context.WithCancel(context.Background())
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	orderService := getOrderService(ctx, conf)

	wg.Add(1)
	go runKafkaConsumer(ctx, &wg, &conf.KafkaConsumer, orderService)

	app := fiber.New()
	config_fiber_app(app, orderService)

	go func() {
		<-shutdown
		log.Println("Shutting down gracefully...")
		cancel()           // Останавливаем Consumer
		_ = app.Shutdown() // Останавливаем HTTP-сервер
	}()

	log.Println("Starting HTTP server on :3000...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}

	wg.Wait()
	log.Println("Server stopped.")
}
