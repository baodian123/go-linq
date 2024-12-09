package main

import (
	"testing"

	"github.com/baodian123/go-linq"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "Basic union",
			first:    []int{1, 2, 3},
			second:   []int{3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty first slice",
			first:    []int{},
			second:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty second slice",
			first:    []int{1, 2, 3},
			second:   []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Duplicates in same slice",
			first:    []int{1, 1, 2, 2},
			second:   []int{3, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := linq.Union(tt.first, tt.second)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
			}
			for i, v := range tt.expected {
				if result[i] != v {
					t.Errorf("at index %d, expected %d, got %d", i, v, result[i])
				}
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "Basic intersection",
			first:    []int{1, 2, 3},
			second:   []int{2, 3, 4},
			expected: []int{2, 3},
		},
		{
			name:     "No intersection",
			first:    []int{1, 2},
			second:   []int{3, 4},
			expected: []int{},
		},
		{
			name:     "Empty first slice",
			first:    []int{},
			second:   []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Duplicates in slices",
			first:    []int{1, 1, 2, 2},
			second:   []int{2, 2, 3},
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := linq.Intersect(tt.first, tt.second)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
			}
			for i, v := range tt.expected {
				if result[i] != v {
					t.Errorf("at index %d, expected %d, got %d", i, v, result[i])
				}
			}
		})
	}
}

func TestExcept(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "Basic except",
			first:    []int{1, 2, 3},
			second:   []int{2, 3, 4},
			expected: []int{1},
		},
		{
			name:     "No common elements",
			first:    []int{1, 2},
			second:   []int{3, 4},
			expected: []int{1, 2},
		},
		{
			name:     "Empty first slice",
			first:    []int{},
			second:   []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Empty second slice",
			first:    []int{1, 2, 3},
			second:   []int{},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := linq.Except(tt.first, tt.second)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
			}
			for i, v := range tt.expected {
				if result[i] != v {
					t.Errorf("at index %d, expected %d, got %d", i, v, result[i])
				}
			}
		})
	}
}

func TestSymmetricExcept(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "Basic symmetric except",
			first:    []int{1, 2, 3},
			second:   []int{2, 3, 4},
			expected: []int{1, 4},
		},
		{
			name:     "No common elements",
			first:    []int{1, 2},
			second:   []int{3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Empty first slice",
			first:    []int{},
			second:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Identical slices",
			first:    []int{1, 2, 3},
			second:   []int{1, 2, 3},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := linq.SymmetricExcept(tt.first, tt.second)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
			}
			for i, v := range tt.expected {
				if result[i] != v {
					t.Errorf("at index %d, expected %d, got %d", i, v, result[i])
				}
			}
		})
	}
}
