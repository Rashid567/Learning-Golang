/*
Разворот строки
Разработать программу, которая переворачивает подаваемую на вход строку.

Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).
*/

package main

import (
	"fmt"
	"slices"
)

// Не подходит для составных символов, так как нарушается структура. К примеру: 1️⃣2️⃣3️⃣
func ReverseString(value string) string {
	runes := []rune(value)
	slices.Reverse(runes)
	return string(runes)
}

func main() {

	strings2Reverse := []string{
		"главрыба",
		"👍👋😊",
		"1️⃣2️⃣3️⃣",
	}

	for _, str := range strings2Reverse {
		reversedStr := ReverseString(str)
		fmt.Println("Original:", str)
		fmt.Println("Reversed:", reversedStr)
		fmt.Println()
	}

}
