/*
Большие числа и операции

Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b,
значения которых > 2^20 (больше 1 миллион).

Комментарий: в Go тип int справится с такими числами, но обратите внимание на возможное переполнение для ещё больших значений.
Для очень больших чисел можно использовать math/big.
*/

package main

import (
	"fmt"
	"math/big"
	"strings"
)

func MakeOperations(a, b *big.Int) {
	fmt.Printf("a: %s\n", a.String())
	fmt.Printf("b: %s\n", b.String())

	addRes := new(big.Int).Add(a, b)
	subRes := new(big.Int).Sub(a, b)
	mulRes := new(big.Int).Mul(a, b)
	divRes := new(big.Int).Div(a, b)

	fmt.Printf("+: %s\n", addRes.String())
	fmt.Printf("-: %s\n", subRes.String())
	fmt.Printf("*: %s\n", mulRes.String())
	fmt.Printf("/: %s\n", divRes.String())
}

func main() {
	a := big.Int{}
	b := big.Int{}

	a.SetString(strings.Repeat("9", 66), 10)
	b.SetString(strings.Repeat("1", 66), 10)

	MakeOperations(&a, &b)
}
