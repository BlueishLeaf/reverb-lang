package evaluator

import (
	"fmt"

	"github.com/BlueishLeaf/go-sound/output"
	"github.com/BlueishLeaf/go-sound/sounds"
	"github.com/BlueishLeaf/go-sound/utilities"

	"github.com/BlueishLeaf/reverb-lang/object"
)

var builtins = map[string]*object.Builtin{
	"length": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 1)
		}
		switch arg := args[0].(type) {
		case *object.Array:
			return &object.Integer{Value: int64(len(arg.Elements))}
		case *object.String:
			return &object.Integer{Value: int64(len(arg.Value))}
		default:
			return newError("argument to `length` not supported, got %s", args[0].Type())
		}
	}},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return Null
		},
	},
	"head": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 1)
		}
		if args[0].Type() != object.ArrayObj {
			return newError("argument to `head` must be ARRAY, got %s", args[0].Type())
		}
		arr := args[0].(*object.Array)
		if len(arr.Elements) > 0 {
			return arr.Elements[0]
		}
		return Null
	}},
	"tail": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 1)
		}
		if args[0].Type() != object.ArrayObj {
			return newError("argument to `tail` must be ARRAY, got %s", args[0].Type())
		}
		arr := args[0].(*object.Array)
		length := len(arr.Elements)
		if length > 0 {
			newElements := make([]object.Object, length-1, length-1)
			copy(newElements, arr.Elements[1:length])
			return &object.Array{Elements: newElements}
		}
		return Null
	}},
	"push": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.ArrayObj {
			return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
		}
		arr := args[0].(*object.Array)
		length := len(arr.Elements)
		newElements := make([]object.Object, length+1, length+1)
		copy(newElements, arr.Elements)
		newElements[length] = args[1]
		return &object.Array{Elements: newElements}
	}},
	"sine": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.FloatObj {
			return newError("argument 0 to `sine` must be FLOAT, got %s", args[0].Type())
		}
		freq := args[0].(*object.Float)
		if args[1].Type() != object.IntegerObj {
			return newError("argument 1 to `sine` must be INTEGER, got %s", args[1].Type())
		}
		duration := args[1].(*object.Integer)
		return &object.Sound{Value: sounds.NewTimedSound(sounds.NewSineWave(freq.Value), float64(duration.Value))}
	}},
	"square": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.FloatObj {
			return newError("argument 0 to `square` must be FLOAT, got %s", args[0].Type())
		}
		freq := args[0].(*object.Float)
		if args[1].Type() != object.IntegerObj {
			return newError("argument 1 to `square` must be INTEGER, got %s", args[1].Type())
		}
		duration := args[1].(*object.Integer)
		return &object.Sound{Value: sounds.NewTimedSound(sounds.NewSquareWave(freq.Value), float64(duration.Value))}
	}},
	"sawtooth": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.FloatObj {
			return newError("argument 0 to `sawtooth` must be FLOAT, got %s", args[0].Type())
		}
		freq := args[0].(*object.Float)
		if args[1].Type() != object.IntegerObj {
			return newError("argument 1 to `sawtooth` must be INTEGER, got %s", args[1].Type())
		}
		duration := args[1].(*object.Integer)
		return &object.Sound{Value: sounds.NewTimedSound(sounds.NewSawtoothWave(freq.Value), float64(duration.Value))}
	}},
	"triangle": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.FloatObj {
			return newError("argument 0 to `triangle` must be FLOAT, got %s", args[0].Type())
		}
		freq := args[0].(*object.Float)
		if args[1].Type() != object.IntegerObj {
			return newError("argument 1 to `triangle` must be INTEGER, got %s", args[1].Type())
		}
		duration := args[1].(*object.Integer)
		return &object.Sound{Value: sounds.NewTimedSound(sounds.NewTriangleWave(freq.Value), float64(duration.Value))}
	}},
	"string": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 3 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.FloatObj {
			return newError("argument 0 to `string` must be FLOAT, got %s", args[0].Type())
		}
		freq := args[0].(*object.Float)
		if args[1].Type() != object.FloatObj {
			return newError("argument 1 to `string` must be FLOAT, got %s", args[0].Type())
		}
		sustain := args[1].(*object.Float)
		if args[2].Type() != object.IntegerObj {
			return newError("argument 2 to `string` must be INTEGER, got %s", args[1].Type())
		}
		duration := args[2].(*object.Integer)
		return &object.Sound{Value: sounds.NewTimedSound(sounds.NewKarplusStrong(freq.Value, sustain.Value), float64(duration.Value))}
	}},
	"silence": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 1)
		}
		if args[0].Type() != object.IntegerObj {
			return newError("argument to `silence` must be INTEGER, got %s", args[0].Type())
		}
		duration := args[0].(*object.Integer)
		return &object.Sound{Value: sounds.NewTimedSilence(float64(duration.Value))}
	}},
	"midiToHz": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 1)
		}
		if args[0].Type() != object.IntegerObj {
			return newError("argument to `midiToHz` must be INTEGER, got %s", args[0].Type())
		}
		midi := args[0].(*object.Integer)
		return &object.Float{Value: utilities.MidiToHz(int(midi.Value))}
	}},
	"concatinate": {Fn: func(args ...object.Object) object.Object {
		if len(args) == 1 {
			if args[0].Type() != object.ArrayObj {
				return newError("argument 1 to `concatinate` must be ARRAY, got %s", args[0].Type())
			}
			arg := args[0].(*object.Array)
			var toConcat []sounds.Sound
			for _, obj := range arg.Elements {
				toConcat = append(toConcat, obj.(*object.Sound).Value)
			}
			return &object.Sound{Value: sounds.ConcatSounds(toConcat...)}
		} else if len(args) > 1 {
			var toConcat []sounds.Sound
			for _, obj := range args {
				toConcat = append(toConcat, obj.(*object.Sound).Value)
			}
			return &object.Sound{Value: sounds.ConcatSounds(toConcat...)}
		}
		return newError("wrong number of arguments. got=%d, want=1 or more", len(args))
	}},
	"overlay": {Fn: func(args ...object.Object) object.Object {
		if len(args) == 1 {
			if args[0].Type() != object.ArrayObj {
				return newError("argument 1 to `overlay` must be ARRAY, got %s", args[0].Type())
			}
			arg := args[0].(*object.Array)
			var toSum []sounds.Sound
			for _, obj := range arg.Elements {
				toSum = append(toSum, obj.(*object.Sound).Value)
			}
			return &object.Sound{Value: sounds.SumSounds(toSum...)}
		} else if len(args) > 1 {
			var toSum []sounds.Sound
			for _, obj := range args {
				toSum = append(toSum, obj.(*object.Sound).Value)
			}
			return &object.Sound{Value: sounds.SumSounds(toSum...)}
		}
		return newError("wrong number of arguments. got=%d, want=1 or more", len(args))
	}},
	"play": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 1)
		}
		switch arg := args[0].(type) {
		case *object.Sound:
			output.Play(arg.Value)
			return Null
		case *object.Array:
			var toPlay []sounds.Sound
			for _, obj := range arg.Elements {
				toPlay = append(toPlay, obj.(*object.Sound).Value)
			}
			output.Play(sounds.ConcatSounds(toPlay...))
			return Null
		default:
			return newError("argument to `play` not supported, got %s", args[0].Type())
		}
	}},
}
