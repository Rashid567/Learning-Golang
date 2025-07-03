/*
Встраивание структур
Дана структура Human (с произвольным набором полей и методов).

Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     uint8
}

func (h *Human) Greet(name string) {
	fmt.Printf("I greet you, %s!\n", name)
}

func (h *Human) Drive(car string) {
	fmt.Printf("%s stars driving %s\n", h.Name, car)
}

type Action struct {
	Human
}

func (a *Action) Greet() {
	fmt.Println("Hello there!")
}

func main() {
	kenobi := Human{
		Name:    "Obi-Wan",
		Surname: "Kenobi",
		Age:     uint8(42),
	}

	action := Action{
		kenobi,
	}

	action.Greet()
	action.Human.Greet("General Grievous")
	action.Human.Drive("Lizard")
}
