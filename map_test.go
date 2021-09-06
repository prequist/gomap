package gomap_test

import (
	"github.com/prequist/gomap"
	"github.com/prequist/ptypes"
	"testing"
)

// Test a plain list.
func TestPlain(t *testing.T) {
	// Values.
	v1 := 1
	v2 := 2
	v3 := 3
	values := []int{v1, v2, v3}
	arr := Plain(v1, v2, v3)
	// Indexing.
	for index, v := range arr {
		// Checking.
		if v != values[index]+1 {
			t.Errorf("received %d, expected %d", v, values[index]+1)
		}
	}
}

// Benchmark the plain list.
func BenchmarkPlain(b *testing.B) {
	// Values.
	v1 := 1
	v2 := 2
	v3 := 3
	values := []int{v1, v2, v3}
	// Start benchmarking.
	for i := 0; i < b.N; i++ {
		// Create an array.
		arr := Plain(v1, v2, v3)
		// Iterate through.
		for index, v := range arr {
			if v != values[index]+1 {
				b.Errorf("received %d, expected %d", v, values[index]+1)
			}
		}
		// Start again.
	}
}

// Test the boxed list.
func TestBox(t *testing.T) {
	// Values.
	v1 := 1
	v2 := 2
	v3 := 3
	values := []int{v1, v2, v3}
	arr := Boxed(v1, v2, v3)
	// Indexing.
	for index, v := range arr {
		// Checking.
		if v != values[index]+2 {
			t.Errorf("received %d, expected %d", v, values[index]+2)
		}
	}
}

// Benchmark the boxed list.
func BenchmarkBox(b *testing.B) {
	// Values.
	v1 := 1
	v2 := 2
	v3 := 3
	values := []int{v1, v2, v3}
	// Benchmarking.
	for i := 0; i < b.N; i++ {
		// Create a new list.
		arr := Boxed(v1, v2, v3)
		// Indexing.
		for index, v := range arr {
			// Checking.
			if v != values[index]+2 {
				b.Errorf("received %d, expected %d", v, values[index]+2)
			}
		}
	}
}

// Test a string list.
func BenchmarkString(b *testing.B) {
	// Values.
	v1 := "hello"
	v2 := "world"
	v3 := "test"
	values := []string{v1, v2, v3}
	// Benchmarking.
	for i := 0; i < b.N; i++ {
		// Create a list.
		arr := String(v1, v2, v3)
		for index, v := range arr {
			// Checking.
			if len(v) != len(values[index])+3 {
				b.Errorf("received size %d expected %d", len(v), len(values[index])+2)
			}
		}
	}
}

// Benchmark a boxed string list.
func BenchmarkBoxedString(b *testing.B) {
	// Values.
	v1 := "hello"
	v2 := "world"
	v3 := "test"
	values := []string{v1, v2, v3}
	// Benchmarking.
	for i := 0; i < b.N; i++ {
		// Create an array.
		arr := BoxedString(v1, v2, v3)
		// Checking.
		for index, v := range arr {
			if len(v) != len(values[index])+3 {
				b.Errorf("received size %d expected %d", len(v), len(values[index])+2)
			}
		}
	}
}

// Benchmark the add function.
func BenchmarkAdd(b *testing.B) {
	v1 := "hello"
	v2 := "world"
	v3 := "test"
	arr := gomap.NewBoxed(v1, v2, v3)
	for i := 0; i < b.N; i++ {
		arr.Add("hello")
	}
}

// Create a plain int list.
func Plain(v ...interface{}) []int {
	e := gomap.New(v...)
	mappable := e.Mappable()
	e = mappable.Map(func(v interface{}) interface{} {
		return v.(int) + 1
	})
	arr := make([]int, len(e.Items()))
	for index, vi := range e.Items() {
		arr[index] = vi.(int)
	}
	return arr
}

// Create a plain string list.
func String(v ...interface{}) []string {
	e := gomap.New(v...)
	mappable := e.Mappable()
	e = mappable.Map(func(v interface{}) interface{} {
		return v.(string) + "---"
	})
	arr := make([]string, len(e.Items()))
	for index, vi := range e.Items() {
		arr[index] = vi.(string)
	}
	return arr
}

// Create a boxed int list.
func Boxed(v ...interface{}) []int {
	e := gomap.NewBoxed(v...)
	mappable := e.Mappable()
	e = mappable.Map(func(v interface{}) interface{} {
		ib := v.(ptypes.Box).IntBox()
		return ptypes.FromInt(*ib.Int() + 2)
	})
	arr := make([]int, len(e.Items()))
	for index, vi := range e.Items() {
		arr[index] = *vi.(ptypes.Box).IntBox().Int()
	}
	return arr
}

// Create a boxed string list.
func BoxedString(v ...interface{}) []string {
	e := gomap.NewBoxed(v...)
	mappable := e.Mappable()
	e = mappable.Map(func(v interface{}) interface{} {
		box := v.(ptypes.Box)
		ib, _ := box.String()
		return ptypes.FromString(ib + "---")
	})
	arr := make([]string, len(e.Items()))
	for index, vi := range e.Items() {
		box := vi.(ptypes.Box)
		str, _ := box.String()
		arr[index] = str
	}
	return arr
}
