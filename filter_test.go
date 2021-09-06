package gomap_test

import (
	"github.com/prequist/gomap"
	"reflect"
	"strings"
	"testing"
)

// Test the filtering function MappableList.Filter(Predicate).
func TestFilter(t *testing.T) {
	stringMap := gomap.New("test", "test1", "test2", "not_applicable", "test3")
	mappable := stringMap.Mappable()
	filtered := mappable.Filter(func(i interface{}) bool {
		if str, ok := i.(string); ok {
			return strings.Contains(str, "test")
		}
		return false
	})
	filteredItems := mapToString(filtered.Items())
	knownKeys := []string{"test", "test1", "test2", "test3"}
	if !reflect.DeepEqual(filteredItems, knownKeys) {
		t.Fail()
	}
}

// Benchmark the filtering function.
func BenchmarkFilter(b *testing.B) {
	stringMap := gomap.New("test", "test1", "test2", "not_applicable", "test3")
	mappable := stringMap.Mappable()
	// Benchmark the actual filtering.
	b.Run("Filter", func(b *testing.B) {
		mappable = mappable.Filter(func(i interface{}) bool {
			if str, ok := i.(string); ok {
				return strings.Contains(str, "test")
			}
			return false
		})
	})
	filteredItems := mapToString(mappable.Items())
	knownKeys := []string{"test", "test1", "test2", "test3"}
	if !reflect.DeepEqual(filteredItems, knownKeys) {
		b.Fail()
	}
}

// Map an interface array to a string.
func mapToString(i []interface{}) []string {
	var strArr []string
	for _, v :=  range i {
		strArr = append(strArr, v.(string))
	}
	return strArr
}