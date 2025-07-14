package main

import (
	"testing"
)

var testFunctions = []struct {
	version  string
	function func([]int, int) int
}{
	{"v1", BinarySearchV1[int]},
	{"v2", BinarySearchV2[int]},
	{"v3", BinarySearchV3[int]},
}

var testCases = []struct {
	name   string
	arr    []int
	target int
	want   int
}{
	{"element exists", []int{1, 3, 5, 7, 9}, 5, 2},
	{"first element", []int{1, 3, 5, 7, 9}, 1, 0},
	{"last element", []int{1, 3, 5, 7, 9}, 9, 4},
	{"element not exists", []int{1, 3, 5, 7, 9}, 4, -1},
	{"empty array", []int{}, 5, -1},
}

func TestBinarySearch(t *testing.T) {

	for _, testFunc := range testFunctions {
		t.Run(
			testFunc.version,
			func(t *testing.T) {
				for _, testCase := range testCases {
					t.Run(
						testCase.name,
						func(t *testing.T) {
							if got := testFunc.function(testCase.arr, testCase.target); got != testCase.want {
								t.Errorf("BinarySearch%s() = %v, want %v", testFunc.version, got, testCase.want)
							}
						},
					)
				}
			},
		)
	}

}
