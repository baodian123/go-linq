package linq

func First[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, item := range slice {
		if predicate(item) {
			return item, true
		}
	}

	var zeroValue T
	return zeroValue, false
}

func FirstOrDefault[T any](slice []T, predicate func(T) bool) T {
	for _, item := range slice {
		if predicate(item) {
			return item
		}
	}

	var zeroValue T
	return zeroValue
}
