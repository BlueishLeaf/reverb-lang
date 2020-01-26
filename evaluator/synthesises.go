package evaluator

import (
	"github.com/BlueishLeaf/reverb-lang/object"
	"github.com/BlueishLeaf/reverb-lang/synth"
	"github.com/hajimehoshi/oto"
	"sync"
)

//var queue = make()

var synthesises = map[string]*object.Synthesis{
	"sine": {Fn: func(ctx *oto.Context, args ...object.Object) object.Object {
		if ctx == nil {
			return newError("oto context is nil, bad news bud")
		}
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.FLOAT_OBJ {
			return newError("argument to `push` must be FLOAT_OBJ, got %s", args[0].Type())
		}
		freq := args[0].(*object.Float)
		duration := args[1].(*object.Integer)
		// TODO: Create a new sine wave and add it to the global queue
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := synth.Play(ctx, freq.Value, duration.Value*1000); err != nil {
				panic(err)
			}
		}()
		wg.Wait()
		return NULL
	}},
}
