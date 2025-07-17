/*
Уникальные символы в строке
Разработать программу, которая проверяет, что все символы в строке
встречаются один раз (т.е. строка состоит из уникальных символов).

Вывод:
true, если все символы уникальны,
false, если есть повторения.
Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/

package main

import (
	"strings"
)

func IsSymbolsUnique(str string) bool {
	seen := map[rune]bool{}

	for _, symbol := range strings.ToLower(str) {
		if _, ok := seen[symbol]; ok {
			return false
		}
		seen[symbol] = true
	}
	return true
}
