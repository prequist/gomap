package gomap

import (
	"sync"
)

// New creates a new list with the given elements.
func New(v ...interface{}) *List {
	// Create the list and return.
	return &List{elements: v}
}

// The list primitive, private type alias.
// listPrimitive is used as an alias for the
// `interface{}` array, handling all elements.
type listPrimitive []interface{}

// The List struct is the actual list.
type List struct {
	// The concurrency lock, mu.
	mu sync.Mutex
	// The list elements.
	elements listPrimitive
	// The list's iterator.
	// The iterator should be accessed via the function
	// Iterator().
	iterator *Iterator
}

// Items returns the items in the list.
func (list *List) Items() []interface{} {
	return list.elements
}

// Add an item to the list.
func (list *List) Add(v interface{}) *List {
	// Lock the mutex.
	list.mu.Lock()
	// Append the elements to the list and the iterator.
	list.elements = append(list.elements, v)
	it := list.Iterator()
	*it.elements = append(*it.elements, v)
	// Unlock the mutex.
	list.mu.Unlock()
	return list
}

// The Remove function removes the item at the requested index.
func (list *List) Remove(index int) *List {
	// If the length of the elements is within the index.
	if len(list.elements) <= index && index > -1 {
		// Lock the mutex.
		list.mu.Lock()
		// Set the element at the index to nil.
		list.elements[index] = nil
		// Unlock the mutex.
		defer list.mu.Unlock()
	}
	// Return the list.
	return list
}

// Iterator returns the list's iterator. If it does not exist,
// it will create a new one.
func (list *List) Iterator() *Iterator {
	// If the iterator does not exist,
	// we can create a new one.
	if list.iterator == nil {
		list.iterator = &Iterator{
			// Add the list's elements.
			elements: (*[]interface{})(&list.elements),
			// Set the iterator index.
			current: -1,
			// Set the current iteration value.
			Value: nil,
		}
	}
	// Return the list's iterator.
	return list.iterator
}
