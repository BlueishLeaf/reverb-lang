package object

// Environment represents a virtual environment within Reverb
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get returns an object within the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set assigns a new object within the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// NewEnvironment creates a new environment instance
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates an extended environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
