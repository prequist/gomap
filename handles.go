package gomap

// The Handler is a type for a slice handler function.
type Handler func(i interface{}, args ...interface{}) (*List, error)

// handleString handles strings.
func handleString(slice interface{}, args ...interface{}) (*List, error) {
	// Arguments are ignored, so we can just remap
	// the slice.
	remapped := remappers["string"]
	return New(remap(&remapped, slice)...), nil
}

// handleInt handles integers.
func handleInt(slice interface{}, args ...interface{}) (*List, error) {
	// Check for arguments.
	if args == nil {
		// If they don't exist, don't pass any int.
		return doHandleInt(slice, nil)
	}
	// Pass in the width.
	width := args[0].(int)
	return doHandleInt(slice, &width)
}

// handleUint handles unsigned integers.
func handleUint(slice interface{}, args ...interface{}) (*List, error) {
	// Check for arguments.
	if args == nil {
		// If they don't exist, don't pass any int.
		return doHandleUint(slice, nil)
	}
	// Pass in the width.
	width := args[0].(int)
	return doHandleUint(slice, &width)
}

// handleFloat handles floats.
func handleFloat(slice interface{}, args ...interface{}) (*List, error) {
	// Check for arguments.
	if args == nil {
		// If they don't exist, don't pass any int.
		return doHandleFloat(slice, nil)
	}
	// Pass in the width.
	width := args[0].(int)
	return doHandleFloat(slice, &width)
}

// handleComplex handles complex numbers.
func handleComplex(slice interface{}, args ...interface{}) (*List, error) {
	// Check for arguments.
	if args == nil {
		// If they don't exist, don't pass any int.
		return doHandleComplex(slice, nil)
	}
	// Pass in the width.
	width := args[0].(int)
	return doHandleComplex(slice, &width)
}

// handleUintptr handles uintptrs.
func handleUintptr(slice interface{}, args ...interface{}) (*List, error) {
	// Arguments are ignored, so we can just remap
	// the slice.
	remapped := remappers["uintptr"]
	return New(remap(&remapped, slice)...), nil
}

// doHandleInt does the actual handling for the integer handle.
func doHandleInt(slice interface{}, width *int) (*List, error) {
	// Do the actual handling.
	remapped := remappers["int"]
	// Get the argument.
	arg := decide(width == nil, nil, *width)
	// Create the new list.
	return New(remap(&remapped, slice, arg)...), nil
}

// doHandleUint does the actual handling for the unsigned integer handle.
func doHandleUint(slice interface{}, width *int) (*List, error) {
	// Do the actual handling.
	remapped := remappers["uint"]
	// Get the argument.
	arg := decide(*width == 0, nil, *width)
	// Create the new list.
	return New(remap(&remapped, slice, arg)...), nil
}

// doHandleFloat does the actual handling for the float handle.
func doHandleFloat(slice interface{}, width *int) (*List, error) {
	// Do the actual handling.
	remapped := remappers["float"]
	// Get the argument.
	arg := decide(width == nil, nil, *width)
	// Create the new list.
	return New(remap(&remapped, slice, arg)...), nil
}

// doHandleComplex does the actual handling for the complex handle.
func doHandleComplex(slice interface{}, width *int) (*List, error) {
	// Do the actual handling.
	remapped := remappers["complex"]
	// Get the argument.
	arg := decide(width == nil, nil, *width)
	// Create the new list.
	return New(remap(&remapped, slice, arg)...), nil
}

// The function decide is a short hand for a ternary type invocation.
func decide(condition bool, optionA interface{}, optionB interface{}) interface{} {
	// If the condition is true,
	// return option a.
	if condition {
		return optionA
	}
	// Otherwise, return option b.
	return optionB
}
