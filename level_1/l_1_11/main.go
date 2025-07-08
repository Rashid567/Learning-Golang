/*
Пересечение множеств

Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы,
присутствующие и в первом, и во втором.

Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}
*/

package main

import (
	"fmt"
)

// O(n * m)  AND  O(1)
func getIntersectionV1[T comparable](array_1 []T, array_2 []T) (res []T) {
	for _, i := range array_1 {
		for _, j := range array_2 {
			if i == j {
				res = append(res, j)
				break
			}
		}
	}
	return
}

// O(n + m)  AND  O(n)
func getIntersectionV2[T comparable](array_1 []T, array_2 []T) (res []T) {
	array_1_as_map := map[T]bool{}
	for _, i := range array_1 {
		array_1_as_map[i] = true
	}

	for _, j := range array_2 {
		if array_1_as_map[j] {
			res = append(res, j)
		}
	}
	return
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	fmt.Println(getIntersectionV1(A, B))
	fmt.Println(getIntersectionV2(A, B))
}
