package linq

func Select[T any, U any](slice []T, selector func(T) U) []U {
	result := make([]U, 0)

	for _, item := range slice {
		result = append(result, selector(item))
	}

	return result
}
