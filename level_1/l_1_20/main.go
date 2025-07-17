/*
Разворот слов в предложении
Разработать программу, которая переворачивает порядок слов в строке.

Пример:
входная строка: «snow dog sun»
выход: «sun dog snow».

Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы,
а выполнять операцию «на месте».
*/

package main

import (
	"fmt"
	"slices"
)

func ReverseSentence(value string) string {
	runes := []rune(value)
	slices.Reverse(runes)

	for start := 0; start < len(runes); {
		// Ищем начало слова - первый не пустой символ
		if runes[start] == ' ' {
			start++
			continue
		}

		// Ищем конец слова - первый пустой символ
		end := start
		for i := end; i < len(runes); i++ {
			if runes[i] == ' ' {
				end = i - 1
				break
			}
		}

		// Переворачиваем слово
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		// Сдвигаем start
		start = end + 1

	}

	return string(runes)

}

func main() {

	sentences := []string{
		"snow dog sun",
		"",
		"  2",
		"    snow   dog   sun  ",
		"    snow   dog   sun  👍  👋😊 ",
	}

	for _, sentence := range sentences {
		reversed := ReverseSentence(sentence)
		fmt.Printf("Original (%d): '%s'\n", len(sentence), sentence)
		fmt.Printf("Reversed (%d): '%s'\n", len(reversed), reversed)
		fmt.Println()
	}

}
