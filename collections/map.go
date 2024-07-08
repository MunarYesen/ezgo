// Package collections provides a set of functions to work with collections in Go.
package collections

func MapBy[T any, F any](sources []T, convert func(source T) F) []F {
	resultSlice := make([]F, len(sources))
	for i, source := range sources {
		resultSlice[i] = convert(source)
	}
	return resultSlice
}

func AsMap[T any, F comparable](sources []T, convert func(source T) F) map[F]T {
	resultMap := make(map[F]T)
	for _, source := range sources {
		resultMap[convert(source)] = source
	}
	return resultMap
}

func GroupBy[T any, F comparable](sources []T, convert func(source T) F) map[F][]T {
	resultMap := make(map[F][]T)
	for _, source := range sources {
		resultMap[convert(source)] = append(resultMap[convert(source)], source)
	}
	return resultMap
}
