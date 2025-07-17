/*
Своя функция Sleep
Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep,
которая приостанавливает выполнение текущей горутины.

Важно: в отличии от настоящей time.Sleep, ваша функция должна именно блокировать выполнение
(например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.

Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).
*/

package main

import (
	"time"
)

// Циклы
func sleepV1(duration time.Duration) {
	startedAt := time.Now()
	for time.Since(startedAt) < duration {
		continue
	}
}

// Таймер
func sleepV2(duration time.Duration) {
	<-time.After(duration)
}

// Через канал
func sleepV3(duration time.Duration) {
	ch := make(chan bool)

	go func() {
		time.Sleep(duration)
		close(ch)
	}()

	<-ch
}
