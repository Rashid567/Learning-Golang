/*
Конкурентный счетчик
Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
По завершению программы структура должна выводить итоговое значение счётчика.

Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterV1 struct {
	value int64
}

func (c *CounterV1) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *CounterV1) Get() int64 {
	return c.value
}

type CounterV2 struct {
	mu    sync.RWMutex
	value int64
}

func (c *CounterV2) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *CounterV2) Get() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	counterV1 := CounterV1{}
	counterV2 := CounterV2{}
	workersCount := 100
	iterationsCount := 100

	wg := sync.WaitGroup{}
	wg.Add(workersCount)

	for range workersCount {
		go func() {
			defer wg.Done()
			for range iterationsCount {
				counterV1.Increment()
				counterV2.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counterV1.Get())
	fmt.Println(counterV2.Get())
}
