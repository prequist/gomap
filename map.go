package gomap

// The Transformer is a type alias for
// a mapping function.
type Transformer func(v interface{}) interface{}

// A MappableList is able to apply
// a Transformer to it's elements.
type MappableList struct {
	// The embedded list.
	*List
}

// The Map function is called on a MappableList to apply
// the requested Transformer.
func (l *MappableList) Map(transformer Transformer) *List {
	// Convert into the parent object.
	return Map(l.List, transformer)
}

// The Map function handles mapping for a raw List.
func Map(list *List, transformer Transformer) *List {
	// Get the iterator.
	itr := list.Iterator()
	// Utilise the iterator recursively, to map
	// the entire list.
	if next, value := itr.Next(); next {
		// Transform the item.
		list.Items()[itr.current] = transformer(value)
		// Remap everything.
		return Map(list, transformer)
	}
	// Return the list.
	return list
}
