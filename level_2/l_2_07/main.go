/*
Что выведет программа?

Объяснить работу конвейера с использованием select.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil {
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Print(v)
	}
}

/*
ОТВЕТ

Вывод программы:
Числа 1-8 в случайном порядке.
Так как sleep вызывается после отправки числа в канал, то обычно первые два числа равны 1 и 2.


Объяснить работу конвейера с использованием select:
select возвращает первое доступное значение. Если их несколько, то в случайном порядке.
Здесь он ждёт первый доступный результат из двух каналов (строка 30).
Прежде чем положить число в результирующий канал идёт проверка закрытости канала.
Если канал закрыт, то переопределяет его значения на nil,
что бы прекратить выполнение функции ожидания и закрыть результирующий канал, если оба канала закрыты (строка 44).

А в строке 58, мы итерируемся по результирующему каналу и распечатываем значения (в range есть обработка закрытости канала)
*/
