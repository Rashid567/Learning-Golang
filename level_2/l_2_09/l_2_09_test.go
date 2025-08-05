package l_2_09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var okTestCases = []struct {
	name     string
	value    string
	validRes string
}{
	{"T1", "a4bc2d5e", "aaaabccddddde"},
	{"T2", "abcd", "abcd"},
	{"T3", "", ""},
	{"T4", "qwe\\4\\5", "qwe45"},
	{"T5", "qwe\\45", "qwe44444"},
}

var errTestCases = []struct {
	name     string
	value    string
	validErr string
}{
	{"T1", "45", "invalid input: rune #0 is digit"},
	{"T2", "qwe\\4\\", "invalid input: \\ in the end"},
}

func TestOk(t *testing.T) {
	for _, tc := range okTestCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := UnpackString(tc.value)
			assert.Equal(t, res, tc.validRes)
			assert.Equal(t, err, nil)
		})
	}
}

func TestErr(t *testing.T) {
	for _, tc := range errTestCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := UnpackString(tc.value)
			assert.Equal(t, res, "")
			assert.Equal(t, err.Error(), tc.validErr)
		})
	}
}
