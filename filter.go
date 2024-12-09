package linq

func Where[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)

	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Distinct[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, item := range slice {
		if _, exists := seen[item]; !exists {
			result = append(result, item)
			seen[item] = struct{}{}
		}
	}

	return result
}
