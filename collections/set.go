package collections

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() *Set[T] {
	s := make(Set[T])
	return &s
}

func (s *Set[T]) Add(element T) bool {
	if _, ok := (*s)[element]; ok {
		return false
	}
	(*s)[element] = struct{}{}
	return true
}

func (s *Set[T]) Remove(element T) bool {
	if _, ok := (*s)[element]; !ok {
		return false
	}
	delete(*s, element)
	return true
}

func (s *Set[T]) Length() int {
	return len(*s)
}
