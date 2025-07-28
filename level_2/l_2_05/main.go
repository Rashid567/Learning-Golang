/*
Что выведет программа?

Объяснить вывод программы.
*/
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

/*
ОТВЕТ

Вывод программы:
```
error
```

Объяснение:
Задача подобная задаче l_2_03, но с двумя отличиями:
- Реализуем собственную структуру ошибки, а не используем готовую
- интерфейс объявляем в основной Gorutine, а в функции возвращаем нулевой указатель

Причина поведения программы такая же.
У интерфейса заполняется тип (itab._type) и методы (itab.fun), а вот значение (data) остаётся nil.
А при сравнении учитывается и тип и значение. Тип != nil и как следствие err != nil

*/
