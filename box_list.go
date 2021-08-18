package gomap

import (
	"github.com/prequist/ptypes"
)

// NewBoxed creates a new boxed List.
func NewBoxed(v ...interface{}) *List {
	// Return a list with the arguments handled.
	return New(handle(v...)...)
}

// The handle function handles any variadic argument and
// turns it into a boxed interface.
func handle(v ...interface{}) []interface{} {
	// Create the new array.
	var arr []interface{}
	// Iterate.
	for _, value := range v {
		// Append the boxed type.
		arr = append(arr, ptypes.FromInterface(value))
	}
	// Return the new array.
	return arr
}
