/*
Обмен значениями без третьей переменной
Поменять местами два числа без использования временной переменной.

Подсказка: примените сложение/вычитание или XOR-обмен.
*/

package main

import (
	"fmt"
)

// Сложение и вычитание
func v1(A, B int) {
	fmt.Println("+-")
	fmt.Println("Before: ", A, B)
	A = A + B
	B = A - B
	A = A - B
	fmt.Println("After:  ", A, B)
}

// XOR
func v2(A, B int) {
	fmt.Println("XOR")
	fmt.Println("Before: ", A, B)
	A = A ^ B
	B = B ^ A
	A = A ^ B
	fmt.Println("After:  ", A, B)
}

func main() {
	A := 7
	B := 8
	v1(A, B)
	v2(A, B)
}
