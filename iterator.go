package gomap

// The Iterator is a small struct
// used for handling iterator-like usage.
// Go itself does not support the while loop,
// however, we can achieve similar usage recursively.
type Iterator struct {
	// The elements.
	elements *[]interface{}
	// The current index.
	current int
	// The current value.
	Value interface{}
}

// Next gets the next value in the iterator
func (it *Iterator) Next() (bool, interface{}) {
	if it.elements == nil {
		return false, nil
	}
	// If there are no more values left, return false and nothing.
	if len(*it.elements) <= it.current+1 {
		return false, nil
	}
	// Upgrade the current index.
	it.current = it.current + 1
	// Set the current value.
	it.Value = (*it.elements)[it.current]
	// Return true, and the value.
	return true, it.Value
}
