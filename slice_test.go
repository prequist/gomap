package gomap_test

import (
	"github.com/prequist/gomap"
	"testing"
)

// TestStringHandle tests the slice string handle.
func TestStringHandle(t *testing.T) {
	values := []string{"a", "b", "c"}
	list := gomap.NewFromPrimitiveSlice(values, values[0])
	for index, v := range list.Items() {
		if values[index] != v.(string) {
			t.Errorf("got %s, expected %s", values[index], v.(string))
		}
	}
}

// TestIntHandle tests the int handle.
func TestIntHandle(t *testing.T) {
	values := []int{1, 2, 3}
	list := gomap.NewFromPrimitiveSlice(values, values[0])
	if list == nil {
		t.Error("list was nil")
		return
	}
	for index, v := range list.Items() {
		if values[index] != v.(int) {
			t.Errorf("the value %d was not %d!", values[index], v.(int))
		}
	}
}

// TestInt32Handle tests the int32 handle.
func TestInt32Handle(t *testing.T) {
	values := []int32{1, 2, 3}
	list := gomap.NewFromPrimitiveSlice(values, values[0])
	if list == nil {
		t.Error("list was nil")
		return
	}
	for index, v := range list.Items() {
		if values[index] != v.(int32) {
			t.Errorf("the value %d was not %d!", values[index], v.(int32))
		}
	}
}

// TestUintHandle tests the uint handle.
func TestUintHandle(t *testing.T) {
	// Create a slice.
	values := []uint{1, 2, 3}
	// Create a list.
	list := gomap.NewFromPrimitiveSlice(values, values[0])
	// If the list is nil, error.
	if list == nil {
		t.Error("list was nil")
		return
	}
	for index, v := range list.Items() {
		if values[index] != v.(uint) {
			t.Errorf("the value %d was not %d!", values[index], v.(uint))
		}
	}
}

// TestUintHandleMap tests mapping for the uint handle.
func TestUintHandleMap(t *testing.T) {
	transformer := func(i interface{}) interface{} {
		return i.(uint) * 2
	}
	values := []uint{1, 3, 5}
	list := gomap.NewFromPrimitiveSlice(values, values[0])
	for index, v := range list.Items() {
		if v != values[index] {
			t.Errorf("got %d, expected %d", v, values[index])
		}
	}
	mappable := list.Mappable()
	list = mappable.Map(transformer)
	for index, v := range list.Items() {
		if v != values[index] * 2 {
			t.Errorf("got %d, expected %d", v, values[index] * 2)
		}
	}
}

// TestFloatHandle tests the float handle.
func TestFloatHandle(t *testing.T) {
	values := []float32{1.2222, 2.3333, 3.4444}
	list := gomap.NewFromPrimitiveSlice(values, values[0])
	if list == nil {
		t.Error("list was nil")
		return
	}
	for index, v := range list.Items() {
		if values[index] != v.(float32) {
			t.Errorf("the value %f was not %f!", values[index], v.(float32))
		}
	}
}

// TestComplexHandle tests the complex handle.
func TestComplexHandle(t *testing.T) {
	values := []complex64{1, 2, 3}
	list := gomap.NewFromPrimitiveSlice(values, values[0])
	if list == nil {
		t.Error("list was nil")
		return
	}
	for index, v := range list.Items() {
		if values[index] != v.(complex64) {
			t.Errorf("the value %f was not %f!", values[index], v.(complex64))
		}
	}
}

// TestComplexHandle tests the complex handle.
func TestUintptrHandle(t *testing.T) {
	values := []uintptr{1, 2, 3}
	list := gomap.NewFromPrimitiveSlice(values, values[0])
	if list == nil {
		t.Error("list was nil")
		return
	}
	for index, v := range list.Items() {
		if values[index] != v.(uintptr) {
			t.Errorf("the value %v was not %v!", values[index], v.(uintptr))
		}
	}
}
