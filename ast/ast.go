package ast

import (
	"bytes"

	"github.com/self-sasi/monkey-interpreter/token"
)

// Base interface for all AST nodes.
// Every node can return the literal value of its associated token.
type Node interface {
	TokenLiteral() string
	String() string
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

func (prog *Program) String() string {
	var out bytes.Buffer

	for _, s := range prog.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Represents a `let` statement in the language, binding a name to a value.
type LetStatement struct {
	Token token.Token // the 'let' token
	Name  *Identifier // identifier being bound (e.g., x, apple)
	Value Expression  // expression assigned to the identifier (e.g., 5, "apple")
}

func (letStatement *LetStatement) statementNode() {}

func (letStatement *LetStatement) TokenLiteral() string {
	return letStatement.Token.Literal
}

func (letStatement *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(letStatement.TokenLiteral() + " ")
	out.WriteString(letStatement.Name.String())
	out.WriteString(" = ")

	if letStatement.Value != nil {
		out.WriteString(letStatement.Value.String())
	} else {
		out.WriteString("nil")
	}

	out.WriteString(";")
	return out.String()
}

// Represents an identifier expression (variable name) like "foo" in "let foo ...".
type Identifier struct {
	Token token.Token // the identifier token (token.IDENT)
	Value string      // the identifier's name ("foo")
}

func (identifier *Identifier) expressionNode() {}

func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }

func (identifier *Identifier) String() string { return identifier.Value }

// Represents a `return` statement in the language, returning an expression.
type ReturnStatement struct {
	Token token.Token // the 'return' token
	Value Expression  // the expression that is returned (e.g., 5, foo)
}

func (returnStatement *ReturnStatement) statementNode() {}

func (returnStatement *ReturnStatement) TokenLiteral() string {
	return returnStatement.Token.Literal
}

func (returnStatement *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(returnStatement.TokenLiteral() + " ")

	if returnStatement.Value != nil {
		out.WriteString(returnStatement.Value.String())
	} else {
		out.WriteString("nil")
	}

	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (expressionStatement *ExpressionStatement) statementNode() {}

func (expressionStatement *ExpressionStatement) TokenLiteral() string {
	return expressionStatement.Token.Literal
}

func (expressionStatement *ExpressionStatement) String() string {
	if expressionStatement.Expression != nil {
		return expressionStatement.Expression.String()
	}
	return ""
}
