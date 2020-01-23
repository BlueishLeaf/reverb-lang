package object

import "github.com/hajimehoshi/oto"

const (
	sampleRate      = 44100
	channelNum      = 2
	bitDepthInBytes = 2
	bufferSize 		= 4096
)

var otoContext, _ = oto.NewContext(sampleRate, channelNum, bitDepthInBytes, bufferSize)

type Environment struct {
	store map[string]Object
	outer *Environment
	otoContext *oto.Context
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e *Environment) GetContext() *oto.Context {
	return e.otoContext
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil, otoContext:otoContext}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
