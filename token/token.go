package token

// Represents the type of a lexical token.
type TokenType string

// Represents a lexical token.
type Token struct {
	Type    TokenType
	Literal string
}

// Token Types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	// Identifiers & literals
	IDENT = "IDENT" // add, foo, bar, x, y, ...
	INT   = "INT"   // integers: 12345...

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

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

// The keywords that exist in the language
var languageKeywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}

// Returns the [TokenType] for identifier if it is a language keyword.
// For example, returns [PLUS] if the identifier is "+".
// If identifier is not a keyword, it returns [IDENT].
func LookupIdentifier(identifier string) TokenType {
	if tokenType, exists := languageKeywords[identifier]; exists {
		return tokenType
	}
	return IDENT
}
