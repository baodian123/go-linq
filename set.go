package linq

func Union[T comparable](first []T, second []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, item := range first {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	for _, item := range second {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

func Intersect[T comparable](first []T, second []T) []T {
	firstSet := make(map[T]struct{})
	result := make([]T, 0)

	for _, item := range first {
		firstSet[item] = struct{}{}
	}

	seen := make(map[T]struct{})
	for _, item := range second {
		if _, exists := firstSet[item]; exists {
			if _, alreadySeen := seen[item]; !alreadySeen {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

func Except[T comparable](first []T, second []T) []T {
	secondSet := make(map[T]struct{})
	for _, item := range second {
		secondSet[item] = struct{}{}
	}

	result := make([]T, 0)
	seen := make(map[T]struct{})
	for _, item := range first {
		if _, existsInSecond := secondSet[item]; !existsInSecond {
			if _, alreadySeen := seen[item]; !alreadySeen {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

func SymmetricExcept[T comparable](first []T, second []T) []T {
	firstSet := make(map[T]struct{})
	secondSet := make(map[T]struct{})

	for _, item := range first {
		firstSet[item] = struct{}{}
	}
	for _, item := range second {
		secondSet[item] = struct{}{}
	}

	result := make([]T, 0)
	seen := make(map[T]struct{})

	for _, item := range first {
		if _, existsInSecond := secondSet[item]; !existsInSecond {
			if _, alreadySeen := seen[item]; !alreadySeen {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	for _, item := range second {
		if _, existsInFirst := firstSet[item]; !existsInFirst {
			if _, alreadySeen := seen[item]; !alreadySeen {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}
