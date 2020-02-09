package evaluator

import (
	"fmt"

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
}
