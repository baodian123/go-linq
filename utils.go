package linq

func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}

	return false
}

func All[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}

	return true
}

func Join[T any, K comparable, U any, R any](left []T, right []U, joinFunc func(T) K, joinKeyFunc func(U) K, mapper func(T, U) R) []R {
	var result []R

	rightMap := make(map[K][]U)
	for _, rightItem := range right {
		key := joinKeyFunc(rightItem)
		rightMap[key] = append(rightMap[key], rightItem)
	}

	for _, leftItem := range left {
		key := joinFunc(leftItem)
		if rightItems, found := rightMap[key]; found {
			for _, rightItem := range rightItems {
				result = append(result, mapper(leftItem, rightItem)) // Return the value of type R
			}
		}
	}

	return result
}
