package main

import (
	"testing"

	"github.com/baodian123/go-linq"
)

func TestFirst(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      int
		wantFound bool
	}{
		{
			name:      "find first even number",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      2,
			wantFound: true,
		},
		{
			name:      "find in empty slice",
			slice:     []int{},
			predicate: func(x int) bool { return true },
			want:      0,
			wantFound: false,
		},
		{
			name:      "no match found",
			slice:     []int{1, 3, 5},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      0,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := linq.First(tt.slice, tt.predicate)
			if got != tt.want || found != tt.wantFound {
				t.Errorf("got = %v, found = %v; want = %v, wantFound = %v", got, found, tt.want, tt.wantFound)
			}
		})
	}

	// Test with string type to verify generic behavior
	strSlice := []string{"apple", "banana", "cherry"}
	strResult, found := linq.First(strSlice, func(s string) bool { return len(s) > 5 })
	if strResult != "banana" || !found {
		t.Errorf("got = %v, found = %v; want = %v, wantFound = %v", strResult, found, "banana", true)
	}
}

func TestFirstOrDefault(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      int
	}{
		{
			name:      "find first even number",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      2,
		},
		{
			name:      "find in empty slice",
			slice:     []int{},
			predicate: func(x int) bool { return true },
			want:      0,
		},
		{
			name:      "no match found",
			slice:     []int{1, 3, 5},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := linq.FirstOrDefault(tt.slice, tt.predicate)
			if got != tt.want {
				t.Errorf("got = %v; want = %v", got, tt.want)
			}
		})
	}

	// Test with string type to verify generic behavior
	strSlice := []string{"apple", "banana", "cherry"}
	strResult := linq.FirstOrDefault(strSlice, func(s string) bool { return len(s) > 5 })
	if strResult != "banana" {
		t.Errorf("got = %v; want = %v", strResult, "banana")
	}
}
