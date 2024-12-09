package main

import (
	"testing"

	"github.com/baodian123/go-linq"
)

func TestWhere(t *testing.T) {
	// Test case 1: Filter even numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := linq.Where(numbers, func(n int) bool {
		return n%2 == 0
	})
	if len(evenNumbers) != 5 {
		t.Errorf("expected 5 even numbers, got %d", len(evenNumbers))
	}
	expectedEvenNumbers := []int{2, 4, 6, 8, 10}
	for i, v := range expectedEvenNumbers {
		if evenNumbers[i] != v {
			t.Errorf("at index %d, expected %d, got %d", i, v, evenNumbers[i])
		}
	}

	// Test case 2: Filter strings with length > 3
	words := []string{"cat", "dog", "elephant", "bird", "rhinoceros"}
	longWords := linq.Where(words, func(w string) bool {
		return len(w) > 3
	})
	if len(longWords) != 3 {
		t.Errorf("expected 3 long words, got %d", len(longWords))
	}
	expectedLongWords := []string{"elephant", "bird", "rhinoceros"}
	for i, v := range expectedLongWords {
		if longWords[i] != v {
			t.Errorf("at index %d, expected %s, got %s", i, v, longWords[i])
		}
	}

	// Test case 3: Empty slice
	emptySlice := []int{}
	result := linq.Where(emptySlice, func(n int) bool {
		return n > 0
	})
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d", len(result))
	}
}

func TestDistinct(t *testing.T) {
	// Test case 1: Integers with duplicates
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	distinctNumbers := linq.Distinct(numbers)
	if len(distinctNumbers) != 5 {
		t.Errorf("expected 5 distinct numbers, got %d", len(distinctNumbers))
	}
	expectedDistinctNumbers := []int{1, 2, 3, 4, 5}
	for i, v := range expectedDistinctNumbers {
		if distinctNumbers[i] != v {
			t.Errorf("at index %d, expected %d, got %d", i, v, distinctNumbers[i])
		}
	}

	// Test case 2: Strings with duplicates
	words := []string{"apple", "banana", "apple", "cherry", "banana", "date"}
	distinctWords := linq.Distinct(words)
	if len(distinctWords) != 4 {
		t.Errorf("expected 4 distinct words, got %d", len(distinctWords))
	}
	expectedDistinctWords := []string{"apple", "banana", "cherry", "date"}
	for i, v := range expectedDistinctWords {
		if distinctWords[i] != v {
			t.Errorf("at index %d, expected %s, got %s", i, v, distinctWords[i])
		}
	}

	// Test case 3: Empty slice
	emptySlice := []int{}
	result := linq.Distinct(emptySlice)
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d", len(result))
	}

	// Test case 4: No duplicates
	uniqueNumbers := []int{1, 2, 3, 4, 5}
	distinctUniqueNumbers := linq.Distinct(uniqueNumbers)
	if len(distinctUniqueNumbers) != 5 {
		t.Errorf("expected 5 distinct unique numbers, got %d", len(distinctUniqueNumbers))
	}
	expectedUniqueNumbers := []int{1, 2, 3, 4, 5}
	for i, v := range expectedUniqueNumbers {
		if distinctUniqueNumbers[i] != v {
			t.Errorf("at index %d, expected %d, got %d", i, v, distinctUniqueNumbers[i])
		}
	}
}
