package gomap

// The Predicate is a type alias for a function that
// takes in a parameter and returns a boolean.
// The predicate is similar to Java's `Predicate<T>`,
// without the generic usage. This makes the usage of a predicate not
// necessarily type safe. This also means that the writer of the predicate
// has to write code that they trust when passing the predicate into the filter.
type Predicate func(interface{}) bool

func (l *MappableList) Filter(predicate Predicate) MappableList {
	// Get the iterator.
	itr := l.Iterator()
	// Utilise the iterator recursively, to map
	// the entire list.
	var newArray List
	// We utilize the goto call to avoid recursively calling the filter function.
sorter:
	if next, value := itr.Next(); next {
		// Filter by the predicate.
		if predicate(value) {
			newArray.Add(value)
		}
		// Jump back to the sorter label.
		goto sorter
	}
	// Return the list.
	return MappableList{&newArray}
}
