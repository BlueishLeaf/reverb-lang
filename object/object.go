package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/BlueishLeaf/reverb-lang/ast"
)

// Type represents the type of an object
type Type string

// BuiltinFunction represent functions that are part of the Reverb standard library
type BuiltinFunction func(args ...Object) Object

const (
	// IntegerObj represents integer objects
	IntegerObj = "INTEGER"
	// FloatObj represents floating point objects
	FloatObj = "FLOAT"
	// BooleanObj represents boolean objects
	BooleanObj = "BOOLEAN"
	// NullObj represents the concept of null
	NullObj = "NULL"
	// ReturnValueObj represents a value returned from a function
	ReturnValueObj = "RETURN_VALUE"
	// ErrorObj represents an error return from the interpreter
	ErrorObj = "ERROR"
	// FunctionObj represents a user-defined function object
	FunctionObj = "FUNCTION"
	// BuiltinObj represents a Reverb function object
	BuiltinObj = "BUILTIN"
	// SynthesisObj represents a Reverb synthesis function object
	SynthesisObj = "SYNTHESIS"
	// ArrayObj represents an array object
	ArrayObj = "ARRAY"
)

// Object represents the smallest component of a Reverb program
type Object interface {
	Type() Type
	Inspect() string
}

// Integer represents the integer data type
type Integer struct {
	Value int64
}

// Type returns the integer object type
func (i *Integer) Type() Type {
	return IntegerObj
}

// Inspect returns the string value of the integer object
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Float represents the floating point data type
type Float struct {
	Value float64
}

// Type returns the float object type
func (f *Float) Type() Type {
	return FloatObj
}

// Inspect returns the string value of the float object
func (f *Float) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}

// Boolean represents the boolean data type
type Boolean struct {
	Value bool
}

// Type returns the boolean object type
func (b *Boolean) Type() Type {
	return BooleanObj
}

// Inspect returns the string value of the boolean object
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null represents the null data type
type Null struct{}

// Type returns the null object type
func (n *Null) Type() Type {
	return NullObj
}

// Inspect returns the string value of the null object
func (n *Null) Inspect() string {
	return "null"
}

// ReturnValue represents the return value object
type ReturnValue struct {
	Value Object
}

// Type returns the return value object type
func (rv *ReturnValue) Type() Type {
	return ReturnValueObj
}

// Inspect returns the string value of the return value object
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error represents an interpreter error
type Error struct {
	Message string
}

// Type returns the error object type
func (e *Error) Type() Type {
	return ErrorObj
}

// Inspect returns the string value of the the error object
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Function represents the function type
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns the function object type
func (f *Function) Type() Type {
	return FunctionObj
}

// Inspect returns the string value of the function object
func (f *Function) Inspect() string {
	var out bytes.Buffer
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString("\n}")
	return out.String()
}

// Builtin represents a builtin function object
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns the builtin function object type
func (b *Builtin) Type() Type {
	return BuiltinObj
}

// Inspect returns the string value of the builtin function object
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// Array represents the array data type
type Array struct {
	Elements []Object
}

// Type returns the array object type
func (a *Array) Type() Type {
	return ArrayObj
}

// Inspect returns the string value of the array object
func (a *Array) Inspect() string {
	var out bytes.Buffer
	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
