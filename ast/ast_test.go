package ast

import (
	"testing"

	"github.com/BlueishLeaf/reverb-lang/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{
					Type:    token.Var,
					Literal: "var",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.Ident,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.Ident,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "var myVar = anotherVar" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
