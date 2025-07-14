package main

import (
	"cmp"
	"fmt"
	"slices"
	"testing"
)

func testQuickSort[T cmp.Ordered](t *testing.T, arrays [][]T) {
	for i, array := range arrays {
		t.Run(
			fmt.Sprintf("Test-%d", i),
			func(t *testing.T) {
				sortedArray := quickSort(array)
				if !slices.IsSorted(sortedArray) {
					t.Errorf("Array (%v) not sorted", sortedArray)
				}
			},
		)
	}
}

func TestQuickSortForInts(t *testing.T) {
	arrays := [][]int{
		{},
		{2},
		{2, 5, 2, -1, 0},
	}
	testQuickSort(t, arrays)
}

func TestQuickSortForFloats(t *testing.T) {
	arrays := [][]float64{
		{},
		{2.0},
		{2.2, 5.3, 2.3, 2.0, -1.0, 0.0},
	}
	testQuickSort(t, arrays)
}

func TestQuickSortForStrings(t *testing.T) {
	arrays := [][]string{
		{},
		{"2"},
		{"2", "5", "2", "-1", "0"},
	}
	testQuickSort(t, arrays)
}
