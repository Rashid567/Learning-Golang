package main


import (
	"fmt"
	"github.com/rashid567/learning-golang/level_2/l_2_09"
)


func main() {
	inputs := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
	}
	for _, input := range inputs {
		res, err := l_2_09.UnpackString(input)
		fmt.Printf("`%s` -> `%s`\n", input, res)
		if err != nil {
			fmt.Printf("Err: `%s`\n", err.Error())
		}
	}
}
