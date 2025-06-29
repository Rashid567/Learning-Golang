package utils

import (
	"fmt"
)

type AppError struct {
	Detail     string
	StatusCode int
}

func (s AppError) Error() string {
	return fmt.Sprintf("%d | %s", s.StatusCode, s.Detail)
}
