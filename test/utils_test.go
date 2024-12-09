package main

import (
	"testing"

	"github.com/baodian123/go-linq"
)

// TestAll tests the All function
func TestAll(t *testing.T) {
	// Test case 1: All elements are even
	numbers := []int{2, 4, 6, 8}
	allEven := linq.All(numbers, func(x int) bool {
		return x%2 == 0
	})
	if allEven != true {
		t.Errorf("expected true, got %v", allEven)
	}

	// Test case 2: Not all elements are even
	mixedNumbers := []int{1, 2, 3, 4}
	allEvenMixed := linq.All(mixedNumbers, func(x int) bool {
		return x%2 == 0
	})
	if allEvenMixed != false {
		t.Errorf("expected false, got %v", allEvenMixed)
	}

	// Test case 3: Empty slice
	emptySlice := []int{}
	allEmpty := linq.All(emptySlice, func(x int) bool {
		return x > 0
	})
	if allEmpty != true { // Conventionally, all elements in an empty slice satisfy the condition
		t.Errorf("expected true, got %v", allEmpty)
	}
}

func TestJoinObjectsWithReturn(t *testing.T) {
	left := []struct {
		ID   int
		Name string
	}{
		{1, "One"},
		{2, "Two"},
		{3, "Three"},
	}

	right := []struct {
		ID   int
		Info string
	}{
		{1, "Info1"},
		{2, "Info2"},
		{3, "Info3"},
	}

	// Define join functions
	joinFunc := func(item struct {
		ID   int
		Name string
	}) int {
		return item.ID
	}
	joinKeyFunc := func(item struct {
		ID   int
		Info string
	}) int {
		return item.ID
	}
	mapper := func(leftItem struct {
		ID   int
		Name string
	}, rightItem struct {
		ID   int
		Info string
	}) struct {
		ID   int
		Name string
		Info string
	} {
		return struct {
			ID   int
			Name string
			Info string
		}{
			ID:   leftItem.ID,
			Name: leftItem.Name,
			Info: rightItem.Info,
		}
	}

	// Perform the join
	result := linq.Join(left, right, joinFunc, joinKeyFunc, mapper)

	// Expected result
	expected := []struct {
		ID   int
		Name string
		Info string
	}{
		{ID: 1, Name: "One", Info: "Info1"},
		{ID: 2, Name: "Two", Info: "Info2"},
		{ID: 3, Name: "Three", Info: "Info3"},
	}

	// Assert the result matches the expected output
	if len(expected) != len(result) {
		t.Errorf("expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("at index %d, expected %+v, got %+v", i, v, result[i])
		}
	}
}
