/*
Конвейер чисел

Разработать конвейер чисел.
Даны два канала:
- в первый пишутся числа x из массива,
- во второй – результат операции x*2.

После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
Убедитесь, что чтение из второго канала корректно завершается.
*/

package main

import (
	"fmt"
	"sync"
)

func RunNumberSquareCalculator(wg *sync.WaitGroup, input <-chan int, output chan<- int) {
	defer wg.Done()
	defer close(output)

	for {
		integer, ok := <-input
		if ok {
			square := integer * integer
			output <- square
		} else {
			break
		}

	}
}

func RunResultsPrinter(wg *sync.WaitGroup, input <-chan int) {
	defer wg.Done()

	for {
		square, ok := <-input
		if ok {
			fmt.Printf("N ^ 2 = %d\n", square)
		} else {
			break
		}
	}
}

func main() {
	inputArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg := sync.WaitGroup{}

	ch_1 := make(chan int)
	ch_2 := make(chan int)

	wg.Add(2)
	go RunNumberSquareCalculator(&wg, ch_1, ch_2)
	go RunResultsPrinter(&wg, ch_2)

	for _, integer := range inputArray {
		ch_1 <- integer
	}
	close(ch_1)

	wg.Wait()
}
