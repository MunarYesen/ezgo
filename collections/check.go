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
