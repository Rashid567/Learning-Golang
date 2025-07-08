/*
Работа нескольких воркеров
Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
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

func GetWorkersCount() uint {
	var workersCount uint
	fmt.Println("Задайте число воркеров (число от 1 до 100 включительно): ")

	for {
		fmt.Scanf("%d", &workersCount)
		if 0 < workersCount && workersCount <= 100 {
			break
		}
		fmt.Println("Невалидное число. Попробуйте снова")
	}

	return workersCount
}

func RunWorker(ctx context.Context, wg *sync.WaitGroup, chanel <-chan string, workerID int) {
	defer wg.Done()

	for {
		select {
		case msg := <-chanel:
			fmt.Printf("Worker #%d got msg: %s\n", workerID, msg)
		case <-ctx.Done():
			fmt.Printf("Worker #%d stopped\n", workerID)
			return
		}
	}
}

func StartWorkers(ctx context.Context, wg *sync.WaitGroup, chanel <-chan string, workersCount uint) {
	wg.Add(int(workersCount))

	for i := 1; i <= int(workersCount); i++ {
		go RunWorker(ctx, wg, chanel, i)
	}
	fmt.Printf("Started %d workers\n", workersCount)
}

func RunProducer(ctx context.Context, chanel chan<- string) {
	defer close(chanel)

	i := 0
	for {
		i++
		msg := fmt.Sprintf("Msg #%d", i)
		select {
		case chanel <- msg:
			time.Sleep(time.Second)
		case <-ctx.Done():
			fmt.Println("Producer stopped")
			return
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg := sync.WaitGroup{}
	ch := make(chan string)
	workersCount := GetWorkersCount()
	StartWorkers(ctx, &wg, ch, workersCount)
	RunProducer(ctx, ch)

	wg.Wait()
}
