package l_2_09

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func UnpackString(input string) (string, error) {
	runes := []rune(input)
	runesLen := len(runes)
	builder := strings.Builder{}

	i := 0
	for i < runesLen {
		// Вытаскиваем руну
		rune := runes[i]
		i++
		if rune == '\\' {
			if i == runesLen {
				return "", fmt.Errorf("invalid input: \\ in the end")
			}
			rune = runes[i]
			i++
		} else if unicode.IsDigit(rune) {
			return "", fmt.Errorf("invalid input: rune #%d is digit", i-1)
		}

		// Определяем кол-во повторений
		digitsCount := 0
		var count int64 = 1
		var err error

		for j := i; j < runesLen; j++ {
			if !unicode.IsDigit(runes[j]) {
				break
			}
			digitsCount++
		}
		if digitsCount != 0 {
			count, err = strconv.ParseInt(string(runes[i:i+digitsCount]), 10, 64)
			if err != nil {
				return "", err
			}
			i += digitsCount
		}

		for range count {
			builder.WriteRune(rune)
		}

	}
	return builder.String(), nil
}
