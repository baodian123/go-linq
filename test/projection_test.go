package main

import (
	"testing"

	"github.com/baodian123/go-linq"
)

func TestSelect(t *testing.T) {
	// Test case 1: Project integers to their squares
	numbers := []int{1, 2, 3, 4, 5}
	squares := linq.Select(numbers, func(x int) int {
		return x * x
	})
	if len(squares) != 5 || squares[0] != 1 || squares[1] != 4 || squares[2] != 9 || squares[3] != 16 || squares[4] != 25 {
		t.Errorf("Expected squares to be [1, 4, 9, 16, 25], got %v", squares)
	}

	// Test case 2: Project strings to their lengths
	strings := []string{"a", "ab", "abc", "abcd"}
	lengths := linq.Select(strings, func(s string) int {
		return len(s)
	})
	if len(lengths) != 4 || lengths[0] != 1 || lengths[1] != 2 || lengths[2] != 3 || lengths[3] != 4 {
		t.Errorf("Expected lengths to be [1, 2, 3, 4], got %v", lengths)
	}

	// Test case 3: Test with empty slice
	emptySlice := []int{}
	emptyResult := linq.Select(emptySlice, func(x int) int {
		return x * 2
	})
	if len(emptyResult) != 0 {
		t.Errorf("Expected emptyResult to be [], got %v", emptyResult)
	}

	// Test case 4: Project integers to strings
	nums := []int{1, 2, 3}
	stringified := linq.Select(nums, func(n int) string {
		return string(rune('A' - 1 + n))
	})
	if len(stringified) != 3 || stringified[0] != "A" || stringified[1] != "B" || stringified[2] != "C" {
		t.Errorf("Expected stringified to be [\"A\", \"B\", \"C\"], got %v", stringified)
	}
}
