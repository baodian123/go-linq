package main

import (
	"testing"

	"github.com/baodian123/go-linq"
)

func TestOrderBy(t *testing.T) {
	t.Run("sort integers", func(t *testing.T) {
		// Test data
		input := []int{3, 1, 4, 1, 5, 9, 2, 6}
		expected := []int{1, 1, 2, 3, 4, 5, 6, 9}

		// Sort using identity function as keySelector
		result := linq.OrderBy(input, func(x int) int { return x }, func(a, b int) bool { return a < b })

		if len(expected) != len(result) {
			t.Errorf("expected length %d, got %d", len(expected), len(result))
		}
		for i, v := range expected {
			if result[i] != v {
				t.Errorf("at index %d, expected %d, got %d", i, v, result[i])
			}
		}
	})

	t.Run("sort strings", func(t *testing.T) {
		// Test data
		input := []string{"banana", "apple", "cherry", "date"}
		expected := []string{"apple", "banana", "cherry", "date"}

		// Sort strings alphabetically
		result := linq.OrderBy(input, func(x string) string { return x }, func(a, b string) bool { return a < b })

		if len(expected) != len(result) {
			t.Errorf("expected length %d, got %d", len(expected), len(result))
		}
		for i, v := range expected {
			if result[i] != v {
				t.Errorf("at index %d, expected %s, got %s", i, v, result[i])
			}
		}
	})

	t.Run("sort structs by field", func(t *testing.T) {
		// Define a test struct
		type Person struct {
			Name string
			Age  int
		}

		// Test data
		input := []Person{
			{Name: "Bob", Age: 30},
			{Name: "Alice", Age: 25},
			{Name: "Charlie", Age: 35},
		}
		expected := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 35},
		}

		// Sort by Name
		result := linq.OrderBy(input, func(p Person) string { return p.Name }, func(a, b string) bool { return a < b })

		if len(expected) != len(result) {
			t.Errorf("expected length %d, got %d", len(expected), len(result))
		}
		for i, v := range expected {
			if result[i] != v {
				t.Errorf("at index %d, expected %+v, got %+v", i, v, result[i])
			}
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var input []int
		result := linq.OrderBy(input, func(x int) int { return x }, func(a, b int) bool { return a < b })
		if len(result) != 0 {
			t.Errorf("expected empty result, got %v", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{1}
		expected := []int{1}
		result := linq.OrderBy(input, func(x int) int { return x }, func(a, b int) bool { return a < b })
		if len(expected) != len(result) {
			t.Errorf("expected length %d, got %d", len(expected), len(result))
		}
		for i, v := range expected {
			if result[i] != v {
				t.Errorf("at index %d, expected %d, got %d", i, v, result[i])
			}
		}
	})
}
