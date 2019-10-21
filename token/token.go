package token

const (
	// Meta
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers & Literals
	IDENT = "IDENT"
	INT   = "INT"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	// Delimiters
	NEWLINE  = "\n" // TODO: May need to also add carriage return '\r' for windows compatibility
	COMMA    = ","
	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"
	// Keywords
	BEGIN    = "BEGIN"
	END      = "END"
	THEN     = "THEN"
	DO       = "DO"
	FOR      = "FOR"
	TO       = "TO"
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Place all keywords for Reverb here
var keywords = map[string]TokenType{
	"begin":  BEGIN,
	"end":    END,
	"echo":   FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"then": THEN,
	"else":   ELSE,
	"do":     DO,
	"for":    FOR,
	"to":     TO,
	"return": RETURN,
}

func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
