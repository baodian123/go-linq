package linq

func GroupBy[T any, K comparable](slice []T, keySelector func(T) K) map[K][]T {
	result := make(map[K][]T)

	for _, item := range slice {
		key := keySelector(item)
		result[key] = append(result[key], item)
	}

	return result
}
