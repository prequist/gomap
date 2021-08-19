package gomap

// The RemapHandler is the function type for
// any handler for remapping.
type RemapHandler func(slice interface{}, args ...interface{}) []interface{}

// remapString remaps a string.
func remapString(slice interface{}, args ...interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]string)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapInt remaps an integer.
func remapInt(slice interface{}, args ...interface{}) []interface{} {
	// Check if there are arguments.
	if args != nil {
		// Switch through the arguments and
		// call the correct remap function.
		width := args[0].(int)
		switch width {
		case 8:
			return remapInt8(slice)
		case 16:
			return remapInt16(slice)
		case 32:
			return remapInt32(slice)
		case 64:
			return remapInt64(slice)
		}
	}
	// Use the default remap function.
	var translated []interface{}
	asserted := slice.([]int)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapFloat remaps a float.
func remapFloat(slice interface{}, args ...interface{}) []interface{} {
	// Check if there are arguments.
	if args != nil {
		// Switch through the arguments and
		// call the correct remap function.
		width := args[0].(int)
		switch width {
		case 32:
			return remapFloat32(slice)
		case 64:
			return remapFloat64(slice)
		}
	}
	// We should never reach here, however,
	// if we do, we can try to remap it as float64.
	return remapFloat64(slice)
}

// remapComplex remaps a complex number.
func remapComplex(slice interface{}, args ...interface{}) []interface{} {
	// Check if there are arguments.
	if args != nil {
		// Switch through the arguments and
		// call the correct remap function.
		width := args[0].(int)
		switch width {
		case 64:
			return remapComplex64(slice)
		case 128:
			return remapComplex128(slice)
		}
	}
	// We should never reach here, however,
	// if we do, we can try to remap it as float64.
	return remapComplex64(slice)
}


// remapInt8 is the remap function for `int8`.
func remapInt8(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]int8)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapInt16 is the remap function for `int16`.
func remapInt16(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]int16)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapInt32 is the remap function for `int32`.
func remapInt32(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]int32)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapInt64 is the remap function for `int64`.
func remapInt64(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]int64)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapInt remaps an unsigned integer.
func remapUint(slice interface{}, args ...interface{}) []interface{} {
	// Check if there are arguments.
	if args != nil {
		// Switch through the arguments and
		// call the correct remap function.
		width := args[0].(int)
		switch width {
		case 8:
			return remapUint8(slice)
		case 16:
			return remapUint16(slice)
		case 32:
			return remapUint32(slice)
		case 64:
			return remapUint64(slice)
		}
	}
	// Use the default remap function.
	var translated []interface{}
	asserted := slice.([]uint)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapUint8 is the remap function for `uint8`.
func remapUint8(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]uint8)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapUint16 is the remap function for `uint16`.
func remapUint16(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]uint16)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapUint32 is the remap function for `uint32`.
func remapUint32(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]uint32)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapUint64 is the remap function for `uint64`.
func remapUint64(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]uint64)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapFloat32 is the remap function for `float32`.
func remapFloat32(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]float32)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapFloat64 is the remap function for `float64`.
func remapFloat64(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]float64)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapComplex64 is the remap function for `complex64`.
func remapComplex64(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]complex64)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapComplex128 is the remap function for `complex128`.
func remapComplex128(slice interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]complex128)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}

// remapUintptr is the remap function for `uintptr`
func remapUintptr(slice interface{}, args ...interface{}) []interface{} {
	var translated []interface{}
	asserted := slice.([]uintptr)
	for _, v := range asserted {
		translated = append(translated, v)
	}
	return translated
}
