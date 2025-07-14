/*
Бинарный поиск
Реализовать алгоритм бинарного поиска встроенными методами языка.
Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.

Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/

package main

import (
	"cmp"
	"fmt"
)

// Рекурсивный БЕЗ передачи индексов
func BinarySearchV1[T cmp.Ordered](array []T, value T) int {
	if len(array) == 0 {
		return -1
	}

	middleIndex := len(array) / 2
	middleValue := array[middleIndex]

	if middleValue == value {
		return middleIndex

	} else if value < middleValue {
		return BinarySearchV1(array[:middleIndex], value)

	} else {
		res_index := BinarySearchV1(array[middleIndex+1:], value)
		if res_index == -1 {
			return res_index
		} else {
			return middleIndex + 1 + res_index
		}
	}
}

// [Внутренняя реализация] Рекурсивный С передачей индексов
func binarySearchV2[T cmp.Ordered](array []T, value T, leftIndex int, rightIndex int) int {
	if leftIndex > rightIndex {
		return -1
	}

	middleIndex := leftIndex + (rightIndex-leftIndex)/2
	middleValue := array[middleIndex]

	if middleValue == value {
		return middleIndex
	} else if value < middleValue {
		return binarySearchV2(array, value, leftIndex, middleIndex-1)
	} else {
		return binarySearchV2(array, value, middleIndex+1, rightIndex)
	}
}

// [Для клиентов] Рекурсивный С передачей индексов
func BinarySearchV2[T cmp.Ordered](array []T, value T) int {
	return binarySearchV2(array, value, 0, len(array)-1)
}

// Не рекурсивный алгоритм
func BinarySearchV3[T cmp.Ordered](array []T, value T) int {
	leftIndex, rightIndex := 0, len(array)-1

	for {
		if leftIndex > rightIndex {
			return -1
		}

		middleIndex := leftIndex + (rightIndex-leftIndex)/2
		middleValue := array[middleIndex]

		if middleValue == value {
			return middleIndex
		} else if value < middleValue {
			rightIndex = middleIndex - 1
		} else {
			leftIndex = middleIndex + 1
		}

	}

}

func main() {
	array_of_int := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i, v := range array_of_int {
		fmt.Printf("%d has index %d\n", v, i)
		fmt.Println(BinarySearchV1(array_of_int, v))
		fmt.Println(BinarySearchV2(array_of_int, v))
		fmt.Println(BinarySearchV3(array_of_int, v))

	}

}
