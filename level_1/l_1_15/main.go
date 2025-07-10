/*
Небольшой фрагмент кода — проблемы и решение
Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?

Приведите корректный пример реализации.

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
Вопрос: что происходит с переменной justString?
*/

/*
ОТВЕТ

`v := createHugeString(1 << 10)` - Создаёт строку длинной 1024 (если предположить что аргумент - это длина строки)

`justString = v[:100]` - Создаёт срез, который просто ссылает на первые 100 символов строки v

Так как justString является глобальной переменной, то GC не удалит justString и как следствие v

Это приводит к утечке памяти. В данном случае 924 байта/символа



КАК ИСПРАВИТЬ

Вариант 1 - Сразу создаём строку нужной длинны:
	func someFunc() {
		justString = createHugeString(100)
	}


Вариант 2 - Копируем часть строки:

	import strings

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = strings.Clone(v[:100])
	}

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	v := "123123123"
	fmt.Println(strings.Clone(v[:4]))
}
