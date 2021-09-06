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

// Mappable returns a mappable list from an existing List.
func (list *List) Mappable() MappableList {
	return MappableList{list}
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
	var mapped List
sorter:
	if next, value := itr.Next(); next {
		// Transform the item.
		transformed := transformer(value)
		mapped.Add(transformed)
		// Remap everything.
		goto sorter
	}
	// Return the list.
	return &mapped
}
