package gomap

import (
	"testing"
)

func TestIterator(t *testing.T) {
	v1 := "a"
	v2 := "b"
	v3 := "c"
	list := New(v1, v2, v3)
	if ok, value := list.Iterator().Next(); !ok && value != v1 {
		t.Error("the expected argument was not the given argument")
	}
	if ok, value := list.Iterator().Next(); !ok && value != v2 {
		t.Error("the expected argument was not the given argument")
	}
	if ok, value := list.Iterator().Next(); !ok && value != v3 {
		t.Error("the expected argument was not the given argument")
	}
}

func BenchmarkIteratorNext(b *testing.B) {
	v1 := "a"
	v2 := "b"
	v3 := "c"
	list := New(v1, v2, v3)
	for i := 0; i < b.N; i++ {
		list.Iterator().Next()
	}
}
