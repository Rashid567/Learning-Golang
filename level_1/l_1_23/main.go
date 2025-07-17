/*
Удаление элемента слайса
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

Подсказка: можно сдвинуть хвост слайса на место удаляемого
элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.
*/

package main

import (
	"fmt"
)

// Проблемы с указателями. К примеру при множественном вызове
func RemoveElementByIndexV1[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}

	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// Медленный, но надежный и предсказуемый
func RemoveElementByIndexV2[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}

	res := make([]T, len(slice)-1)
	copy(res, slice[:index])
	copy(res[index:], slice[index+1:])
	return res
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("V2 - Медленный, но надежный и предсказуемый")
	for i := range slice {
		fmt.Println(RemoveElementByIndexV2(slice, i))
	}

	fmt.Println()
	fmt.Println("V1 - Проблемы с указателями. К примеру при множественном вызове")
	for i := range slice {
		fmt.Println(RemoveElementByIndexV1(slice, i))
	}
}
