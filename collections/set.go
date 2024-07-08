package collections

type Set[T comparable] struct {
	mapping map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{
		mapping: make(map[T]struct{}),
	}
	return s
}

func (s *Set[T]) Add(element T) bool {
	if _, ok := s.mapping[element]; ok {
		return false
	}
	s.mapping[element] = struct{}{}
	return true
}

func (s *Set[T]) Remove(element T) bool {
	if _, ok := s.mapping[element]; !ok {
		return false
	}
	delete(s.mapping, element)
	return true
}

func (s *Set[T]) Length() int {
	return len(s.mapping)
}

func (s *Set[T]) Has(element T) bool {
	_, ok := s.mapping[element]
	return ok
}

func (s *Set[T]) Slice() []T {
	if s.Length() == 0 {
		return nil
	}
	sl := make([]T, 0, s.Length())
	for el := range s.mapping {
		sl = append(sl, el)
	}
	return sl
}
