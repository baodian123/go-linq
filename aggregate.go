package linq

func Sum[T any, U any](slice []T, accumulator func(U, T) U) U {
	var result U

	for _, item := range slice {
		result = accumulator(result, item)
	}

	return result
}
