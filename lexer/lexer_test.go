package lexer

import (
	"testing"

	"github.com/BlueishLeaf/reverb-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `
var five = 5.56
var ten = 10
var add = fn(x, y):
	x + y
# This comment should be ignored
var result = add(five, ten) # This comment should also be ignored
!-/*5
5 < 10 > 5
if 5 < 10:
	return true
elif 5 == 1:
	return true
else:
	return false
10 == 10
10 != 9
[1, 2]
%
@
<=
>=
&&
||
&
|
=`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Newline, "\n"},
		{token.Var, "var"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Float, "5.56"},
		{token.Newline, "\n"},
		{token.Var, "var"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Newline, "\n"},
		{token.Var, "var"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LParen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.RParen, ")"},
		{token.Colon, ":"},
		{token.Newline, "\n"},
		{token.Tab, "\t"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Newline, "\n"},
		{token.Comment, "# This comment should be ignored"},
		{token.Newline, "\n"},
		{token.Var, "var"},
		{token.Ident, "result"},
		{token.Assign, "="},
		{token.Ident, "add"},
		{token.LParen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.RParen, ")"},
		{token.Comment, "# This comment should also be ignored"},
		{token.Newline, "\n"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Int, "5"},
		{token.Newline, "\n"},
		{token.Int, "5"},
		{token.LT, "<"},
		{token.Int, "10"},
		{token.GT, ">"},
		{token.Int, "5"},
		{token.Newline, "\n"},
		{token.If, "if"},
		{token.Int, "5"},
		{token.LT, "<"},
		{token.Int, "10"},
		{token.Colon, ":"},
		{token.Newline, "\n"},
		{token.Tab, "\t"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Newline, "\n"},
		{token.Elif, "elif"},
		{token.Int, "5"},
		{token.Equal, "=="},
		{token.Int, "1"},
		{token.Colon, ":"},
		{token.Newline, "\n"},
		{token.Tab, "\t"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Newline, "\n"},
		{token.Else, "else"},
		{token.Colon, ":"},
		{token.Newline, "\n"},
		{token.Tab, "\t"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Newline, "\n"},
		{token.Int, "10"},
		{token.Equal, "=="},
		{token.Int, "10"},
		{token.Newline, "\n"},
		{token.Int, "10"},
		{token.NotEqual, "!="},
		{token.Int, "9"},
		{token.Newline, "\n"},
		{token.LBracket, "["},
		{token.Int, "1"},
		{token.Comma, ","},
		{token.Int, "2"},
		{token.RBracket, "]"},
		{token.Newline, "\n"},
		{token.Mod, "%"},
		{token.Newline, "\n"},
		{token.Illegal, "@"},
		{token.Newline, "\n"},
		{token.LTE, "<="},
		{token.Newline, "\n"},
		{token.GTE, ">="},
		{token.Newline, "\n"},
		{token.And, "&&"},
		{token.Newline, "\n"},
		{token.Or, "||"},
		{token.Newline, "\n"},
		{token.Illegal, "&"},
		{token.Newline, "\n"},
		{token.Illegal, "|"},
		{token.Newline, "\n"},
		{token.Assign, "="},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q, literal=%q", i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q, tokentype=%q", i, tt.expectedLiteral, tok.Literal, tok.Type)
		}
	}
}
