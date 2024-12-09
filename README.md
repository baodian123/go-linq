# Go LINQ

A LINQ-inspired library for Go that provides a set of extension methods for working with slices and collections.

## LINQ Functions

### Projection
- `Select[T, U](slice []T, selector func(T) U) []U` - Projects each element of a slice into a new form.

### Sorting
- `OrderBy[T, K](slice []T, keySelector func(T) K, less func(K, K) bool) []T` - Sorts elements in ascending order based on a key.

### Set Operations
- `Union[T](first []T, second []T) []T` - Produces the set union of two slices.
- `Intersect[T](first []T, second []T) []T` - Produces the set intersection of two slices.
- `Except[T](first []T, second []T) []T` - Produces the set difference of two slices.
- `SymmetricExcept[T](first []T, second []T) []T` - Produces the symmetric difference of two slices.

### Grouping
- `GroupBy[T, K](slice []T, keySelector func(T) K) map[K][]T` - Groups elements by a specified key.

### Filtering
- `Any[T](slice []T, predicate func(T) bool) bool` - Tests if any element satisfies a condition.
- `All[T](slice []T, predicate func(T) bool) bool` - Tests if all elements satisfy a condition.
- `Where[T any](slice []T, predicate func(T) bool) []T` - Filters elements based on a predicate.
- `Distinct[T comparable](slice []T) []T` - Returns a slice with distinct elements.

### Aggregation
- `Sum[T any, U any](slice []T, accumulator func(U, T) U) U` - Computes the sum of elements in a slice using an accumulator function.

### Element Access
- `First[T any](slice []T, predicate func(T) bool) (T, bool)` - Returns the first element that satisfies the predicate and a boolean indicating success.
- `FirstOrDefault[T any](slice []T, predicate func(T) bool) T` - Returns the first element that satisfies the predicate or a default value.

### Join Operations
- `Join[T any, K comparable, U any, R any](left []T, right []U, joinFunc func(T) K, joinKeyFunc func(U) K, mapper func(T, U) R) []R` - Joins two slices based on a key, using the provided join functions and returns a slice of the mapped results.

## Sample Usage

Here are some examples of how to use the LINQ-inspired functions in this library.

### Union

```go
package main

import (
    "fmt"
    "github.com/baodian123/go-linq"
)

func main() {
    first := []int{1, 2, 3}
    second := []int{3, 4, 5}
    result := linq.Union(first, second)
    fmt.Println(result) // Output: [1 2 3 4 5]
}
```

### Intersect

```go
package main

import (
    "fmt"
    "github.com/baodian123/go-linq"
)

func main() {
    first := []int{1, 2, 3, 4}
    second := []int{3, 4, 5, 6}
    result := linq.Intersect(first, second)
    fmt.Println(result) // Output: [3 4]
}
```

### GroupBy

```go
package main

import (
    "fmt"
    "github.com/baodian123/go-linq"
)

func main() {
    items := []struct {
        Name  string
        Group string
    }{
        {"Amily", "A"},
        {"Bob", "B"},
        {"David", "A"},
    }
    result := linq.GroupBy(items, func(item struct {
        Name  string
        Group string
    }) string {
        return item.Group
    })
    fmt.Println(result) // Output: map[A:[{Amily A} {David A}] B:[{Bob B}]]
}
```

### Sum

```go
package main

import (
    "fmt"
    "github.com/baodian123/go-linq"
)

func main() {
    numbers := []int{1, 2, 3, 4}
    result := linq.Sum(numbers, func(acc int, item int) int {
        return acc + item
    })
    fmt.Println(result) // Output: 10
}
```

### Where

```go
package main

import (
    "fmt"
    "github.com/baodian123/go-linq"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    result := linq.Where(numbers, func(item int) bool {
        return item%2 == 0
    })
    fmt.Println(result) // Output: [2 4]
}
```

### Why No `From` Syntax?

In this Go LINQ library, we intentionally chose not to include the `From` syntax, which is common in other LINQ implementations like in C#. This decision was made to align with Go's design philosophy, which emphasizes simplicity, clarity, and minimalism.

Go is known for its straightforward, explicit style, and we wanted the library to feel natural for Go developers. Instead of using `From` as a query syntax starter (similar to SQL-like or LINQ queries), we rely on clear, chainable function calls like `Where`, `Select`, and `GroupBy`. This approach is consistent with Go's idiomatic way of handling operations on collections, where you compose operations using functions in a straightforward, readable manner.

While the `From` syntax may be more familiar to users coming from languages like C# or JavaScript, using a more function-based approach keeps the library lightweight, intuitive, and easier to integrate with Go's existing ecosystem.

By keeping the API simple and function-based, we ensure the library remains consistent with Go's core principles and minimizes unnecessary abstraction.