package evaluator

import (
	"github.com/BlueishLeaf/reverb-lang/lexer"
	"github.com/BlueishLeaf/reverb-lang/object"
	"github.com/BlueishLeaf/reverb-lang/parser"
	"testing"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer, got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, expected=%d", result.Value, expected)
		return false
	}
	return true
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float, got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, expected=%f", result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean, got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, expected=%t", result.Value, expected)
		return false
	}
	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.56", 5.56},
		{"10.104", 10.104},
		{"-5.56", -5.56},
		{"-10.104", -10.104},
		{"5.2 + 5.2 + 5.2 + 5.2 + 5.2 - 10", 16},
		{"2.1 * 2.69 * 2.44 * 2.22 * 2.1", 64.25895672},
		{"-50.5 + 100.5 + -50.5", -0.5},
		{"5.5 * 2 + 10", 21},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestEvalBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestEvalIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if true then 10 end", 10},
		{"if false then 10 end", nil},
		{"if 1 then 10 end", 10},
		{"if 1 < 2 then 10 end", 10},
		{"if 1 > 2 then 10 end", nil},
		{"if 1 > 2 then 10 end else then 20 end", 20},
		{"if 1 < 2 then 10 end else then 20 end", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestEvalReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10", 10},
		{"return 2 * 5", 10},
		{"if 10 > 1 then return 10 end", 10},
		{
			`if 10 > 1 then
				  if 10 > 1 then
					return 10
				  end
				
				  return 1
				end`,
			10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"true + false + true + false",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if 10 > 1 then true + false end",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`999[1]`,
			"index operator not supported: INTEGER",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestEvalVarStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var a = 5\na", 5},
		{"var a = 5 * 5\na", 25},
		{"var a = 5\nvar b = a\nb", 5},
		{"var a = 5\nvar b = a\nvar c = a + b + 5\nc", 15},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestEvalFunctionObject(t *testing.T) {
	input := "echo(x) begin x + 2 end"
	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}
	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}
	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}
	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var identity = echo(x) begin x end\nidentity(5)", 5},
		{"var identity = echo(x) begin return x end\nidentity(5)", 5},
		{"var double = echo(x) begin x * 2 end\ndouble(5)", 10},
		{"var add = echo(x, y) begin x + y end\nadd(5, 5)", 10},
		{"var add = echo(x, y) begin x + y end\nadd(5 + 5, add(5, 5))", 20},
		{"echo(x) begin x end (5)", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestEnclosingEnvironments(t *testing.T) {
	input := `var first = 10
		var second = 10
		var third = 10
		
		var ourFunction = echo(first) begin
		  var second = 20
		
		  first + second + third
		end
		
		ourFunction(20) + first + second`

	testIntegerObject(t, testEval(input), 70)
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len([1, 2, 3])`, 3},
		{`len([])`, 0},
		{`first([1, 2, 3])`, 1},
		{`first([])`, nil},
		{`first(1)`, "argument to `first` must be ARRAY, got INTEGER"},
		{`last([1, 2, 3])`, 3},
		{`last([])`, nil},
		{`last(1)`, "argument to `last` must be ARRAY, got INTEGER"},
		{`rest([1, 2, 3])`, []int{2, 3}},
		{`rest([])`, nil},
		{`push([], 1)`, []int{1}},
		{`push(1, 1)`, "argument to `push` must be ARRAY, got INTEGER"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case nil:
			testNullObject(t, evaluated)
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)",
					evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q",
					expected, errObj.Message)
			}
			case []int:
				array, ok := evaluated.(*object.Array)
				if !ok {
					t.Errorf("obj not Array. got=%T (%+v)", evaluated, evaluated)
					continue
				}

				if len(array.Elements) != len(expected) {
					t.Errorf("wrong num of elements. want=%d, got=%d",
						len(expected), len(array.Elements))
					continue
				}

				for i, expectedElem := range expected {
					testIntegerObject(t, array.Elements[i], int64(expectedElem))
				}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"var i = 0\n[1][i]",
			1,
		},
		{
			"[1, 2, 3][1 + 1]",
			3,
		},
		{
			"var myArray = [1, 2, 3]\nmyArray[2]",
			3,
		},
		{
			"var myArray = [1, 2, 3]\nmyArray[0] + myArray[1] + myArray[2]",
			6,
		},
		{
			"var myArray = [1, 2, 3]\nvar i = myArray[0]\nmyArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
