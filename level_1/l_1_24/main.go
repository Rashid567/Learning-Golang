/*
Расстояние между точками
Разработать программу нахождения расстояния между двумя точками на плоскости.
Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
Расстояние рассчитывается по формуле между координатами двух точек.

Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(1, 1)
	fmt.Printf("Distance between Point%+v and Point%+v = %f\n", p1, p2, p1.Distance(p2))
}
