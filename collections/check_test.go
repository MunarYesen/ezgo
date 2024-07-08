package collections_test

import (
	"munaryesen/ezgo/collections"
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	cases := []struct {
		name     string
		sources  []int
		convert  func(int) bool
		expected bool
	}{
		{"All positive", []int{1, 2, 3}, func(n int) bool { return n > 0 }, true},
		{"Not all positive", []int{-1, 2, 3}, func(n int) bool { return n > 0 }, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := collections.All(tc.sources, tc.convert)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestAny(t *testing.T) {
	cases := []struct {
		name     string
		sources  []int
		convert  func(int) bool
		expected bool
	}{
		{"Any positive", []int{-1, -2, 3}, func(n int) bool { return n > 0 }, true},
		{"No positive", []int{-1, -2, -3}, func(n int) bool { return n > 0 }, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := collections.Any(tc.sources, tc.convert)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	t.Run("TestFirstWithMatchingElement", func(t *testing.T) {
		sources := []int{1, 2, 3, 4}
		result := collections.First(sources, func(n int) bool { return n%2 == 0 })
		if result != 2 {
			t.Errorf("Expected 2, got %v", result)
		}
	})

	t.Run("TestFirstWithNoMatchingElement", func(t *testing.T) {
		sources := []int{1, 3, 5}
		result := collections.First(sources, func(n int) bool { return n%2 == 0 })
		var expected int
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("TestFilterWithMatchingElements", func(t *testing.T) {
		sources := []int{1, 2, 3, 4}
		result := collections.Filter(sources, func(n int) bool { return n%2 == 0 })
		expected := []int{2, 4}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("TestFilterWithNoMatchingElements", func(t *testing.T) {
		sources := []int{1, 3, 5}
		result := collections.Filter(sources, func(n int) bool { return n%2 == 0 })
		expected := make([]int, 0)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
