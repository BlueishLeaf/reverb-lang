package token

const (
	// Illegal represents an unrecognised symbol
	Illegal = "ILLEGAL"
	// EOF represents the end of the file
	EOF = "EOF"
	// Comment represents a user-defined piece of documentation
	Comment = "#"
	// Ident represents an user-defined identifier
	Ident = "IDENT"
	// Int represents an integer literal
	Int = "INT"
	// Float represents a float literal
	Float = "FLOAT"
	// String represents a string literal
	String = "STRING"
	// Assign represents the assignment operator
	Assign = "="
	// Plus represents the plus operator
	Plus = "+"
	// Minus represents the minus operator
	Minus = "-"
	// Bang represents the bang operator
	Bang = "!"
	// Asterisk represents the asterisk operator
	Asterisk = "*"
	// Slash represents the slash operator
	Slash = "/"
	// Mod represents the modulus operator
	Mod = "%"
	// LT represents the less-than comparison operator
	LT = "<"
	// LTE represents the less-than-or-equal-to comparison operator
	LTE = "<="
	// GT represents the greater-than comparison operator
	GT = ">"
	// GTE represents the greater-than-or-equal-to comparison operator
	GTE = ">="
	// Equal represents the equal-to comparison operator
	Equal = "=="
	// NotEqual represents the not-equal-to comparison operator
	NotEqual = "!="
	// And represents the and operator
	And = "&&"
	// Or represents the or operator
	Or = "||"
	// Newline represents the newline delimiter
	Newline = "\n"
	// Comma represents the comma delimiter
	Comma = ","
	// LParen represents the left parenthesis
	LParen = "("
	// RParen represents the right parenthesis
	RParen = ")"
	// LBracket represents the left bracket
	LBracket = "["
	// RBracket represents the right bracket
	RBracket = "]"
	// Begin represents the beginning of a function block
	Begin = "BEGIN"
	// Then represents the beginning of an If block
	Then = "THEN"
	// End represents the end of a code block
	End = "END"
	// Function represents a function
	Function = "FUNCTION"
	// Var represents a user-defined variable
	Var = "VAR"
	// True represents the boolean symbol for true
	True = "TRUE"
	// False represents the boolean symbol for false
	False = "FALSE"
	// If represents a conditional statement
	If = "IF"
	// Else represents the alternative consequence of a conditional statement
	Else = "ELSE"
	// Return represents the return keyword
	Return = "RETURN"
)

// Type represents the token type
type Type string

// Token represents a source code token
type Token struct {
	Type
	Literal string
}

// LookupKeyword checks an identifier against a map of keyowards and
// returns the corresponding token
func LookupKeyword(keyword string) Type {
	var keywords = map[string]Type{
		"begin":  Begin,
		"end":    End,
		"fn":     Function,
		"var":    Var,
		"true":   True,
		"false":  False,
		"if":     If,
		"then":   Then,
		"else":   Else,
		"return": Return,
	}
	if tok, ok := keywords[keyword]; ok {
		return tok
	}
	return Ident
}
