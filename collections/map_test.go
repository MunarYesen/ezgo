package collections_test

import (
	"github.com/stretchr/testify/assert"
	"munaryesen/ezgo/collections"
	"testing"
)

type testStruct struct {
	fieldInteger int
	fieldString  string
}

func TestMapBy(t *testing.T) {
	sources := []testStruct{{1, "a"}, {2, "b"}, {3, "c"}}
	expected := []int{1, 2, 3}
	notExpected := []int{1, 3, 2}
	result := collections.MapBy(sources, func(source testStruct) int {
		return source.fieldInteger
	})
	assert.Equal(t, result, expected, "should be equal")
	assert.NotEqual(t, result, notExpected, "should be equal")
}

func TestAsMap(t *testing.T) {
	sources := []testStruct{{1, "a"}, {2, "b"}, {3, "c"}}
	expected := map[string]testStruct{
		"a": {1, "a"},
		"b": {2, "b"},
		"c": {3, "c"},
	}
	result := collections.AsMap(sources, func(source testStruct) string {
		return source.fieldString
	})
	assert.Equal(t, result, expected, "should be equal")
}

func TestGroupBy(t *testing.T) {
	sources := []testStruct{
		{1, "a"},
		{1, "a1"},
		{2, "b"},
		{2, "b1"},
		{3, "c"}}
	expected := map[int][]testStruct{
		1: {{1, "a"}, {1, "a1"}},
		2: {{2, "b"}, {2, "b1"}},
		3: {{3, "c"}},
	}
	result := collections.GroupBy(sources, func(source testStruct) int {
		return source.fieldInteger
	})
	assert.Equal(t, result, expected, "should be equal")
}
