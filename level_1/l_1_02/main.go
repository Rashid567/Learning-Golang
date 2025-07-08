/*
Конкурентное возведение в квадрат
Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/

package main

import (
	"fmt"
	"sync"
)

func CalculateSquare(wg *sync.WaitGroup, num int) {
	fmt.Printf("%d^2 = %d\n", num, num*num)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	inputArray := []int{2, 4, 6, 8, 10}
	wg.Add(len(inputArray))

	for _, num := range inputArray {
		go CalculateSquare(&wg, num)
	}

	wg.Wait()
}
