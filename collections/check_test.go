package collections_test

import (
	"munaryesen/ezgo/collections"
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
