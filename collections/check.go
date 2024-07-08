package collections

// All is a function that returns true if all elements in the slice return true when passed to the convert function.
func All[T any](sources []T, convert func(source T) bool) bool {
	for _, source := range sources {
		if !convert(source) {
			return false
		}
	}
	return true
}

// Any is a function that returns true if any element in the slice returns true when passed to the convert function.
func Any[T any](sources []T, convert func(source T) bool) bool {
	for _, source := range sources {
		if convert(source) {
			return true
		}
	}
	return false
}

// First is a function that returns the first element in the slice that meets the convention
func First[T any](sources []T, convert func(source T) bool) T {
	var result T
	for _, source := range sources {
		if convert(source) {
			return source
		}
	}
	return result
}

// Filter is a function that returns a slice containing all elements that meet the convention
func Filter[T any](sources []T, convert func(source T) bool) []T {
	var result = make([]T, 0, len(sources))
	for _, source := range sources {
		if convert(source) {
			result = append(result, source)
		}
	}
	return result
}
