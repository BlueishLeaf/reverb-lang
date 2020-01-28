package evaluator

import (
	"github.com/BlueishLeaf/reverb-lang/object"
	"github.com/BlueishLeaf/reverb-lang/synth"
	"github.com/hajimehoshi/oto"
	"sync"
)

// TODO: Convert to map, with key notations
var pianoSequence = [88]float64{
	27.5,
	29.1353,
	30.8677,
	32.7032,
	34.6479,
	36.7081,
	38.8909,
	41.2035,
	43.6536,
	46.2493,
	48.9995,
	51.9130,
	55.0,
	58.2705,
	61.7345,
	65.4064,
	69.2957,
	73.4162,
	77.7817,
	82.4069,
	87.3071,
	92.4986,
	97.9989,
	103.826,
	110.0,
	116.541,
	123.471,
	130.813,
	138.591,
	146.832,
	155.563,
	164.814,
	174.614,
	184.997,
	195.998,
	207.652,
	220.0,
	233.082,
	246.942,
	261.626,
	277.183,
	293.665,
	311.127,
	329.628,
	349.228,
	369.994,
	391.995,
	415.305,
	440.0, // Yay! halfway there
	466.164,
	493.883,
	523.251,
	554.365,
	587.33,
	622.254,
	659.255,
	698.456,
	739.989,
	783.991,
	830.609,
	880.0,
	932.328,
	987.767,
	1046.5,
	1108.73,
	1174.66,
	1244.51,
	1318.51,
	1396.91,
	1479.98,
	1567.98,
	1661.22,
	1760.0,
	1864.66,
	1975.53,
	2093.0,
	2217.46,
	2349.32,
	2489.02,
	2637.02,
	2793.83,
	2959.96,
	3135.96,
	3322.44,
	3520.0,
	3729.31,
	3951.07,
	4186.01,
}
//var pianoKeys = map[int]float64{
//	0: 27.5,
//
//}

var synthesises = map[string]*object.Synthesis{
	"sine": {Fn: func(ctx *oto.Context, args ...object.Object) object.Object {
		if ctx == nil {
			return newError("oto context is nil, bad news bud")
		}
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.FLOAT_OBJ {
			return newError("argument to `sine` must be FLOAT_OBJ, got %s", args[0].Type())
		}
		if args[1].Type() != object.INTEGER_OBJ {
			return newError("argument to `sine` must be INTEGER_OBJ, got %s", args[0].Type())
		}
		freq := args[0].(*object.Float)
		duration := args[1].(*object.Integer)
		// TODO: Create a new sine wave and add it to the global queue
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := synth.Play(ctx, freq.Value, duration.Value); err != nil {
				panic(err)
			}
		}()
		wg.Wait()
		return NULL
	}},
	"sequence": {Fn: func(ctx *oto.Context, args ...object.Object) object.Object {
		if ctx == nil {
			return newError("oto context is nil, bad news bud")
		}
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=%d", len(args), 2)
		}
		if args[0].Type() != object.INTEGER_OBJ {
			return newError("argument to `sequence` must be INTEGER_OBJ, got %s", args[0].Type())
		}
		if args[1].Type() != object.INTEGER_OBJ {
			return newError("argument to `sequence` must be INTEGER_OBJ, got %s", args[0].Type())
		}
		sequenceTerm := args[0].(*object.Integer)
		duration := args[1].(*object.Integer)
		freq := pianoSequence[sequenceTerm.Value % 88]
		// TODO: Create a new sine wave and add it to the global queue
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := synth.Play(ctx, freq, duration.Value); err != nil {
				panic(err)
			}
		}()
		wg.Wait()
		return NULL
	}},
}
