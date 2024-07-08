package collections_test

import (
	"munaryesen/ezgo/collections"
	"reflect"
	"sort"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := collections.NewSet[int]()
	if s == nil {
		t.Errorf("NewSet() = %v, want non-nil", s)
	}
	if len(*s) != 0 {
		t.Errorf("NewSet() = %d, want 0", len(*s))
	}
}

func TestAdd(t *testing.T) {
	s := collections.NewSet[int]()
	if !s.Add(1) {
		t.Errorf("Add(1) = false, want true")
	}
	if !s.Add(2) {
		t.Errorf("Add(2) = false, want true")
	}
	if s.Add(1) {
		t.Errorf("Add(1) again = true, want false")
	}
}

func TestRemove(t *testing.T) {
	s := collections.NewSet[int]()
	s.Add(1)
	if !s.Remove(1) {
		t.Errorf("Remove(1) = false, want true")
	}
	if s.Remove(2) {
		t.Errorf("Remove(2) = true, want false")
	}
}

func TestLength(t *testing.T) {
	s := collections.NewSet[int]()
	if s.Length() != 0 {
		t.Errorf("Length() = %d, want 0", s.Length())
	}
	s.Add(1)
	s.Add(2)
	if s.Length() != 2 {
		t.Errorf("Length() after adding = %d, want 2", s.Length())
	}
}

func TestSliceNonEmptySet(t *testing.T) {
	s := collections.NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	result := s.Slice()
	if len(result) != 3 {
		t.Errorf("Slice() returned slice of length %d, want 3", len(result))
	}

	expected := []int{1, 2, 3}
	sort.Ints(result) // Sorting to ensure the order does not affect the test
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Slice() = %v, want %v", result, expected)
	}
}

func TestSliceUniqueness(t *testing.T) {
	s := collections.NewSet[int]()
	s.Add(1)
	s.Add(1) // Attempt to add duplicate
	result := s.Slice()
	if len(result) != 1 || result[0] != 1 {
		t.Errorf("Slice() returned %v, want [1] for uniqueness test", result)
	}
}
