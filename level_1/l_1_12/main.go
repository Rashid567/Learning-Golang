/*
Собственное множество строк

Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
*/

package main

import (
	"fmt"
)

// O(n^2)  AND  O(1)
func getUniqueV1[T comparable](array []T) (res []T) {
OuterLoop:
	for _, i := range array {
		for _, j := range res {
			if i == j {
				continue OuterLoop
			}
		}
		res = append(res, i)
	}
	return
}

// O(n + u)  AND  O(u), где u <= n
func getUniqueV2[T comparable](array []T) (res []T) {
	res_as_map := map[T]bool{}
	for _, i := range array {
		if !res_as_map[i] {
			res_as_map[i] = true
		}
	}

	for key := range res_as_map {
		res = append(res, key)
	}

	return
}

func main() {
	A := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(getUniqueV1(A))
	fmt.Println(getUniqueV2(A))
}
