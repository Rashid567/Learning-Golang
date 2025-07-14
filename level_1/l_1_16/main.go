/*
Быстрая сортировка (quicksort)
Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.

Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
Для выбора опорного элемента можно взять середину или первый элемент.
*/

package main

import (
	"cmp"
	"fmt"
)

func quickSort[T cmp.Ordered](array []T) []T {
	if len(array) <= 1 {
		return array
	}
	pivot_index := len(array) / 2
	pivot := array[pivot_index]
	left := []T{}
	right := []T{}

	for i, v := range array {
		if i == pivot_index {
			continue
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(
		append(quickSort(left), pivot),
		quickSort(right)...,
	)

}

func main() {
	unsorted_array_of_int := []int{2, 6, 3, 5, 9, 7, 8, 1, 4}
	unsorted_array_of_str := []string{"2", "6", "3", "5", "9", "7", "8", "1", "4"}

	fmt.Println(quickSort((unsorted_array_of_int)))
	fmt.Println(quickSort((unsorted_array_of_str)))
}
