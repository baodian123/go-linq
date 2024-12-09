package main

import (
	"testing"

	"github.com/baodian123/go-linq"
)

func TestGroupBy(t *testing.T) {
	// Test case 1: Group numbers by their remainder when divided by 3
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	grouped := linq.GroupBy(numbers, func(n int) int {
		return n % 3
	})

	if len(grouped) < 3 {
		t.Errorf("expected at least 3 groups, got %d", len(grouped))
	}
	if len(grouped[0]) != 3 || grouped[0][0] != 3 || grouped[0][1] != 6 || grouped[0][2] != 9 {
		t.Error("Numbers divisible by 3 are incorrect")
	}
	if len(grouped[1]) != 3 || grouped[1][0] != 1 || grouped[1][1] != 4 || grouped[1][2] != 7 {
		t.Error("Numbers with remainder 1 are incorrect")
	}
	if len(grouped[2]) != 3 || grouped[2][0] != 2 || grouped[2][1] != 5 || grouped[2][2] != 8 {
		t.Error("Numbers with remainder 2 are incorrect")
	}

	// Test case 2: Group strings by their length
	words := []string{"cat", "dog", "fish", "bird", "elephant", "lion"}
	groupedWords := linq.GroupBy(words, func(s string) int {
		return len(s)
	})

	if len(groupedWords) < 3 {
		t.Errorf("expected at least 3 groups, got %d", len(groupedWords))
	}
	if len(groupedWords[3]) != 2 || groupedWords[3][0] != "cat" || groupedWords[3][1] != "dog" {
		t.Error("3-letter words are incorrect")
	}
	if len(groupedWords[4]) != 3 || groupedWords[4][0] != "fish" || groupedWords[4][1] != "bird" || groupedWords[4][2] != "lion" {
		t.Error("4-letter words are incorrect")
	}
	if len(groupedWords[8]) != 1 || groupedWords[8][0] != "elephant" {
		t.Error("8-letter words are incorrect")
	}

	// Test case 3: Empty slice
	emptySlice := []int{}
	emptyGrouped := linq.GroupBy(emptySlice, func(n int) int {
		return n % 2
	})

	if len(emptyGrouped) != 0 {
		t.Error("Empty slice should result in empty grouping")
	}
}
