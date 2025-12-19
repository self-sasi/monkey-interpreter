package token

// Represents the type of a lexical token.
type TokenType string

// Represents a lexical token.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	// Identifiers & literals
	IDENT = "IDENT" // add, foo, bar, x, y, ...
	INT   = "INT"   // integers: 12345...

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// Paranthesis & brackets
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
