/*
Завершение по Ctrl+C
Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).

Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func RunWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("Worker still running")
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		}
	}
}

func main() {
	// Обоснование выбора:
	// signal и context - стандартные библиотеки
	// signal.NotifyContext - кратко и удобно.
	// При необходимости можно использовать context.WithTimeout для ограничения времени graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go RunWorker(ctx, &wg)

	wg.Wait()
	fmt.Println("App stopped")
}
