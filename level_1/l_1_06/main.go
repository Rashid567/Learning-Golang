/*
Остановка горутины
Реализовать все возможные способы остановки выполнения горутины.

Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.

Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

// выход по условию
func StopOnCondition(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for i := 0; true; i++ {
		if i >= 3 {
			fmt.Println("StopOnCondition stopped")
			return
		}
		fmt.Println("StopOnCondition still working")
		time.Sleep(time.Second)
	}
}

// выход через канал уведомления
func StopOnChanelNotify(wg *sync.WaitGroup, notifyChanel <-chan bool) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-notifyChanel:
			fmt.Println("StopOnChanelNotify stopped")
			return
		case <-time.After(time.Second):
			fmt.Println("StopOnChanelNotify still working")
		}
	}
}

// выход через контекст
func StopWitchContext(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("StopWitchContext stopped")
			return
		case <-time.After(time.Second):
			fmt.Println("StopWitchContext still working")
		}
	}
}

// выход через прекращение работы runtime.Goexit()
func StopWithGoexit(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for i := 0; true; i++ {
		if i >= 3 {
			fmt.Println("StopWithGoexit stopped")
			runtime.Goexit()
		}
		fmt.Println("StopWithGoexit still working")
		time.Sleep(time.Second)
	}
}

// выход через panic
func StopWithPanic(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in: %s\n", r)
		}
	}()

	for i := 0; true; i++ {
		if i >= 3 {
			panic("StopWithPanic stopped")
		}
		fmt.Println("StopWithPanic still working")
		time.Sleep(time.Second)
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg := sync.WaitGroup{}
	stopChanel := make(chan bool, 1)

	go StopOnCondition(&wg)
	go StopOnChanelNotify(&wg, stopChanel)
	go StopWitchContext(ctx, &wg)
	go StopWithGoexit(&wg)
	go StopWithPanic(&wg)

	time.Sleep(time.Second * 3)
	stopChanel <- true
	stop()

	wg.Wait()
	fmt.Println("App stopped")
}
