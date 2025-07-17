package main

import (
	"testing"
)

var testCases = []struct {
	name     string
	str      string
	validRes bool
}{
	{"T1", "abcd", true},
	{"T2", "abCdefA", false},
	{"T3", "aabcd", false},
}

func Test(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsSymbolsUnique(tc.str)
			if res != tc.validRes {
				t.Errorf("Expected result `%v`. Got `%v`", tc.validRes, res)
			}
		})
	}
}
