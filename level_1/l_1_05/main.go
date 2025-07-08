/*
Таймаут на канал
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

Подсказка: используйте time.After или таймер для ограничения времени работы.
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

// Реализация с помощью time.NewTimer
func RunWriter(ctx context.Context, wg *sync.WaitGroup, ch chan<- int, runTime int) {
	defer wg.Done()
	defer close(ch)

	timer := time.NewTimer(time.Second * time.Duration(runTime))

	fmt.Printf("Writer started (runTime = %d sec)\n", runTime)

	i := 0
	for {
		i++
		select {
		case ch <- i:
			fmt.Printf("Writer sent `%d` to chanel\n", i)
			time.Sleep(time.Millisecond * 100)
		case <-timer.C:
			fmt.Println("Writer completed job")
			return
		case <-ctx.Done():
			fmt.Println("Writer was interrupted")
			return
		}
	}
}

// Реализация с помощью context.WithTimeout
func RunWriterV2(ctx context.Context, wg *sync.WaitGroup, ch chan<- int, runTime int) {
	defer wg.Done()
	defer close(ch)

	stop_ctx, stop := context.WithTimeout(ctx, time.Second*time.Duration(runTime))
	defer stop()

	fmt.Printf("Writer started (runTime = %d sec)\n", runTime)

	i := 0
	for {
		i++
		select {
		case ch <- i:
			fmt.Printf("Writer sent `%d` to chanel\n", i)
			time.Sleep(time.Millisecond * 100)
		case <-stop_ctx.Done():
			if stop_ctx.Err() == context.DeadlineExceeded {
				fmt.Println("Writer completed job")
			} else {
				fmt.Println("Writer was interrupted")
			}
			return
		}
	}
}

func RunReader(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()

	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Printf("Reader got `%d` from chanel\n", i)
			} else {
				fmt.Println("Reader stopped (chanel closed)")
				return
			}

		case <-ctx.Done():
			fmt.Println("Reader stopped (ctx)")
			return
		}
	}
}

func main() {
	runTime := 1

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go RunWriter(ctx, &wg, ch, runTime)
	go RunReader(ctx, &wg, ch)

	wg.Wait()
	fmt.Println("App stopped")
}
