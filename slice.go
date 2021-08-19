package gomap

import (
	"reflect"
	"strconv"
	"strings"
)

// The ContentHandler structure
// is a registration structure
// to hold a handler and a remapper.
type ContentHandler struct {
	// Handler is the handler.
	Handler Handler
	// Mapper is the remapper to `interface{}`.
	Mapper  RemapHandler
}

var (
	// The registered handles.
	handles = make(map[string]Handler, 0)
	// The registered remappers.
	remappers = make(map[string]RemapHandler, 0)

	// StringHandler global instance.
	StringHandler = PutHandler("string", ContentHandler{handleString, remapString})

	// IntHandler global instance.
	IntHandler = PutHandler("int", ContentHandler{handleInt, remapInt})

	// UintHandler global instance.
	UintHandler = PutHandler("uint", ContentHandler{handleUint, remapUint})

	// FloatHandler global instance.
	FloatHandler = PutHandler("float", ContentHandler{handleFloat, remapFloat})

	// ComplexHandler global instance.
	ComplexHandler = PutHandler("complex", ContentHandler{handleComplex, remapComplex})

	// UintptrHandler global instance.
	UintptrHandler = PutHandler("uintptr", ContentHandler{handleUintptr, remapUintptr})
)

// NewFromPrimitiveSlice handles slice support, however, to
// avoid the reflect solution to this,
// we try to handle type assertion ourself.
//
// The parameter element takes the first element.
//
// Note that if the slice is of []interface{},
// the New function may be used instead.
func NewFromPrimitiveSlice(slice interface{}, element interface{}) *List {
	// The type of the first element.
	ty := reflect.TypeOf(element)
	// The name of the type.
	name := ty.Name()
	// The argument to pass.
	var arg interface{}
	// The list that will eventually be returned.
	var list *List

	// setList is the embedded local function to
	// set the data in the list.
	setList := func() {
		// Find the handler.
		handler, ok := handles[name]
		if !ok {
			return
		}
		// Set the list.
		list, _ = handler(slice, arg)
	}

	// If it's a uintptr, exit early.
	if name == "uintptr" {
		setList()
		return list
	}

	// Check if the value is an integer.
	handleNaming(name, "int", func(s []string) {
		name = s[0]
		arg, _ = strconv.Atoi(s[1])
	}).OnFail(
		// Check if the value is a float.
		func() {
			handleNaming(name, "float", func(s []string) {
				name = s[0]
				arg, _ = strconv.Atoi(s[1])
			})
		},
	).OnFail(func() {
		// Check if the value is complex.
		handleNaming(name, "complex", func(s []string) {
			name = s[0]
			arg, _ = strconv.Atoi(s[1])
		})
	}).Cleanup(setList)
	return list
}

// handleNaming checks the a condition and calls the respective callback
// after the string split.
func handleNaming(str string, expected string, callback func(s []string)) Callable {
	// If the string contains the expected string,
	// we can split it, then invoke the callback.
	if strings.Contains(str, expected) {
		split := strings.SplitAfterN(str, expected, 2)
		callback(split)
		return Callable{Passed: true}
	}
	return Callable{Passed: false}
}

// PutHandler puts a handler into the registered handlers map,
// registered by name.
func PutHandler(name string, handler ContentHandler) ContentHandler {
	handles[name] = handler.Handler
	remappers[name] = handler.Mapper
	return handler
}

// The remap function remaps a slice given remapper arguments and a remapper.
func remap(remapper *RemapHandler, slice interface{}, args ...interface{}) []interface{} {
	// If the remapper is nil, we can't go forward.
	if remapper == nil {
		return nil
	}
	// If there are no arguments, we don't have to include them.
	// We can pass only the slice into the invocation.
	if args == nil || args[0] == nil {
		return (*remapper)(slice)
	}
	// There are arguments, so we will spread the arguments and pass
	// them forward.
	return (*remapper)(slice, args...)
}
