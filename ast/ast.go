package ast

import "github.com/self-sasi/monkey-interpreter/token"

// Base interface for all AST nodes.
// Every node can return the literal value of its associated token.
type Node interface {
	TokenLiteral() string
}

// Represents a statement node in the AST
// (e.g., let statements).
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression node in the AST
// (e.g., identifiers, calculations).
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST.
// It consists of a sequence of statements.
type Program struct {
	Statements []Statement
}

// Returns the literal value of the first token in the program,
// or an empty string if the program is empty.
func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) > 0 {
		return prog.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Represents a `let` statement in the language, binding a name to a value.
type LetStatement struct {
	Token token.Token // the 'let' token
	Name  *Identifier // identifier being bound (e.g., x, apple)
	Value Expression  // expression assigned to the identifier (e.g., 5, "apple")
}

func (letStatement *LetStatement) statementNode() {}

func (letStatement *LetStatement) TokenLiteral() string { return letStatement.Token.Literal }

// Represents an identifier expression (variable name) like "foo" in "let foo ...".
type Identifier struct {
	Token token.Token // the identifier token (token.IDENT)
	Value string      // the identifier's name ("foo")
}

func (identifier *Identifier) expressionNode() {}

func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }
