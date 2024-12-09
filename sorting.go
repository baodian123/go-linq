package linq

import "sort"

func OrderBy[T any, K comparable](slice []T, keySelector func(T) K, less func(K, K) bool) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	sort.Slice(result, func(i, j int) bool {
		return less(keySelector(result[i]), keySelector(result[j]))
	})

	return result
}
