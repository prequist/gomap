package gomap

// The Callable structure allows
// chaining of function calls (promise like)
type Callable struct {
	Passed bool
	Result interface{}
}

// Then calls something if the callable has passed.
func (callable Callable) Then(function func(value interface{})) Callable {
	// If the callable has passed we can call the function.
	if callable.Passed {
		function(callable.Result)
	}
	return callable
}

// OnFail calls a function on the failure of a callable
func (callable Callable) OnFail(function func()) Callable {
	// If the callable has failed, we can call the function.
	if !callable.Passed {
		function()
	}
	return callable
}

// Cleanup is the call to end the callable's chain.
func (callable Callable) Cleanup(function func()) {
	// At the end of it's life, we can call the function.
	function()
}


