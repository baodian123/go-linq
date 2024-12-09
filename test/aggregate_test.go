package main

import (
	"testing"

	"github.com/baodian123/go-linq"
)

func TestSum(t *testing.T) {
	t.Run("sum of integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := linq.Sum(numbers, func(acc, curr int) int {
			return acc + curr
		})

		if result != 15 {
			t.Errorf("expected 15, got %d", result)
		}
	})

	t.Run("sum of float64", func(t *testing.T) {
		numbers := []float64{1.5, 2.5, 3.5}
		result := linq.Sum(numbers, func(acc, curr float64) float64 {
			return acc + curr
		})

		if result != 7.5 {
			t.Errorf("expected 7.5, got %f", result)
		}
	})

	t.Run("concatenate strings", func(t *testing.T) {
		words := []string{"Hello", " ", "World"}
		result := linq.Sum(words, func(acc, curr string) string {
			return acc + curr
		})

		if result != "Hello World" {
			t.Errorf("expected 'Hello World', got '%s'", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := linq.Sum(numbers, func(acc, curr int) int {
			return acc + curr
		})

		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})

	t.Run("custom struct accumulation", func(t *testing.T) {
		type Person struct {
			Age int
		}
		people := []Person{{Age: 25}, {Age: 30}, {Age: 35}}
		result := linq.Sum(people, func(acc int, curr Person) int {
			return acc + curr.Age
		})

		if result != 90 {
			t.Errorf("expected 90, got %d", result)
		}
	})
}
