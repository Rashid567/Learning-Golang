/*
Определение типа переменной в runtime
Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать: int, string, bool, chan (канал).

Подсказка: оператор типа switch v.(type) поможет в решении.
*/

package main

import (
	"fmt"
	"reflect"
)

func DetermineValueTypeV1(value any) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		fmt.Printf("V1 | `%d` is integer\n", v)
	case float32, float64:
		fmt.Printf("V1 | `%f` is float\n", v)
	case string:
		fmt.Printf("V1 | `%s` is string\n", v)
	case bool:
		fmt.Printf("V1 | `%t` is bool\n", v)
	case chan int, chan string:
		fmt.Printf("V1 | `%T` is chanel\n", v)
	default:
		fmt.Printf("V1 | Failed to determine type of `%T`\n", v)
	}
}

func DetermineValueTypeV2(value any) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("V2 | `%d` is integer\n", value)
	case reflect.Float32, reflect.Float64:
		fmt.Printf("V2 | `%f` is float\n", value)
	case reflect.String:
		fmt.Printf("V2 | `%s` is string\n", value)
	case reflect.Bool:
		fmt.Printf("V2 | `%t` is bool\n", value)
	case reflect.Chan:
		fmt.Printf("V2 | `%T` is chanel\n", value)
	default:
		fmt.Printf("V2 | Failed to determine type of `%T`\n", value)
	}
}

func main() {

	params := []any{7, 7.0, "7", true, false, make(chan int), main}

	for _, p := range params {
		DetermineValueTypeV1(p)
		DetermineValueTypeV2(p)
	}

}
