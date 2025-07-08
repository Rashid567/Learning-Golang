/*
Конкурентная запись в map
Реализовать безопасную для конкуренции запись данных в структуру map.

Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

Проверьте работу кода на гонки (util go run -race).
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func ConcurrentMapWriterWithMutex() {
	wg := sync.WaitGroup{}
	data := map[string]int{}
	mtx := sync.Mutex{}

	workersCount := 10
	iterationCount := 100
	wg.Add(workersCount)

	for range workersCount {
		go func() {
			defer wg.Done()

			for range iterationCount {
				mtx.Lock()
				data["key"]++
				mtx.Unlock()
			}
		}()
	}

	wg.Wait()

	rightCount := workersCount * iterationCount
	fmt.Printf("rightCount = %d, data[key] = %d\n", rightCount, data["key"])
	if rightCount == data["key"] {
		fmt.Println("rightCount == data[key]")
	} else {
		panic("rightCount != data[key]")
	}
}

func ConcurrentMapWriterWithSyncMap() {
	wg := sync.WaitGroup{}
	data := sync.Map{}

	initValue := int64(1)
	workersCount := 10
	iterationCount := 100
	wg.Add(workersCount)

	for range workersCount {
		go func() {
			defer wg.Done()

			for range iterationCount {
				val, ok := data.LoadOrStore("key", &initValue)
				ptr := val.(*int64)
				if ok {
					atomic.AddInt64(ptr, 1)
				}
			}
		}()
	}

	wg.Wait()

	count, _ := data.Load("key")
	count_ptr := count.(*int64)

	rightCount := int64(workersCount * iterationCount)
	fmt.Printf("rightCount = %d, data[key] = %d\n", rightCount, *count_ptr)
	if rightCount == *count_ptr {
		fmt.Println("rightCount == data[key]")
	} else {
		panic("rightCount != data[key]")
	}
}

func main() {
	fmt.Println("ConcurrentMapWriterWithMutex")
	ConcurrentMapWriterWithMutex()

	fmt.Println("ConcurrentMapWriterWithSyncMap")
	ConcurrentMapWriterWithSyncMap()
}
