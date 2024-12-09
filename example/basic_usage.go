package main

import (
	"fmt"

	"github.com/baodian123/go-linq"
)

func main() {
	// Sample data
	numbers := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	words := []string{"apple", "banana", "apple", "cherry", "date", "banana"}

	fmt.Println("=== Basic Operations ===")

	// Filter operations
	evenNumbers := linq.Where(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evenNumbers) // [2 4 4 2]

	uniqueNumbers := linq.Distinct(numbers)
	fmt.Println("Unique numbers:", uniqueNumbers) // [1 2 3 4 5]

	// Projection operations
	doubled := linq.Select(numbers, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled numbers:", doubled) // [2 4 6 8 10 10 8 6 4 2]

	// Ordering
	ordered := linq.OrderBy(numbers, func(n int) int {
		return n
	}, func(a, b int) bool {
		return a < b
	})
	fmt.Println("Ordered numbers:", ordered) // [1 1 2 2 3 3 4 4 5 5]

	fmt.Println("\n=== Set Operations ===")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	union := linq.Union(slice1, slice2)
	fmt.Println("Union:", union) // [1 2 3 4 5 6 7 8]

	intersect := linq.Intersect(slice1, slice2)
	fmt.Println("Intersect:", intersect) // [4 5]

	except := linq.Except(slice1, slice2)
	fmt.Println("Except:", except) // [1 2 3]

	symmetricExcept := linq.SymmetricExcept(slice1, slice2)
	fmt.Println("Symmetric Except:", symmetricExcept) // [1 2 3 6 7 8]

	fmt.Println("\n=== Element Operations ===")

	// First/FirstOrDefault
	first, _ := linq.First(numbers, func(n int) bool {
		return n > 3
	})
	fmt.Println("First number > 3:", first) // 4

	defaultValue := linq.FirstOrDefault(numbers, func(n int) bool {
		return n > 10
	})
	fmt.Println("First number > 10 (with default):", defaultValue) // -1

	fmt.Println("\n=== Grouping Operations ===")

	// Group words by their first letter
	grouped := linq.GroupBy(words, func(s string) string {
		return string(s[0])
	})
	fmt.Println("Grouped by first letter:", grouped)
	// Output will show groups like:
	// a: [apple apple]
	// b: [banana banana]
	// c: [cherry]
	// d: [date]

	fmt.Println("\n=== Aggregation Operations ===")

	sum := linq.Sum(numbers, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("Sum of numbers:", sum) // 30

	// Chaining operations example
	result := linq.Select(
		linq.Where(numbers, func(n int) bool {
			return n%2 == 0
		}),
		func(n int) int {
			return n * 2
		},
	)
	fmt.Println("\n=== Chaining Operations ===")
	fmt.Println("Even numbers doubled:", result) // [4 8 8 4]

	fmt.Println("\n=== String Operations Example ===")

	fruits := []string{"apple", "banana", "orange", "banana"}
	vegetables := []string{"carrot", "banana", "celery"}

	fmt.Println("Unique fruits:", linq.Distinct(fruits))
	fmt.Println("Common items:", linq.Intersect(fruits, vegetables))

	// Using Any and All
	hasApple := linq.Any(fruits, func(s string) bool {
		return s == "apple"
	})
	allLongWords := linq.All(fruits, func(s string) bool {
		return len(s) > 3
	})

	fmt.Println("Has apple:", hasApple)               // true
	fmt.Println("All words > 3 chars:", allLongWords) // true
}
