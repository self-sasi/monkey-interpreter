package parser

import (
	"fmt"

	"github.com/self-sasi/monkey-interpreter/ast"
	"github.com/self-sasi/monkey-interpreter/lexer"
	"github.com/self-sasi/monkey-interpreter/token"
)

// expression precedence related values
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type (
	prefixParseFunction func() ast.Expression
	infixParseFunction  func(ast.Expression) ast.Expression
)

// Parser transforms a stream of tokens produced by the lexer
// into an Abstract Syntax Tree (AST).
type Parser struct {
	lex       *lexer.Lexer // source of tokens
	curToken  token.Token  // current token under examination
	peekToken token.Token  // next token (one-token lookahead)
	errors    []string     // list of errors

	prefixParseFns map[token.TokenType]prefixParseFunction
	infixParseFns  map[token.TokenType]infixParseFunction
}

// Creates and initializes a new Parser.
func New(lex *lexer.Lexer) *Parser {
	parserPointer := &Parser{
		lex:    lex,
		errors: []string{},
	}

	// read two tokens, so curToken and peekToken are both set
	parserPointer.nextToken()
	parserPointer.nextToken()

	parserPointer.prefixParseFns = make(map[token.TokenType]prefixParseFunction)
	parserPointer.registerPrefix(token.IDENT, parserPointer.parseIdentifier)

	return parserPointer
}

// Advances the parser to the next token.
// The current token becomes the previous peek token,
// and a new peek token is read from the lexer.
func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lex.NextToken()
}

// returns a slice of errors the parser records
func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(expectedToken token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		expectedToken, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFunction) {
	parser.prefixParseFns[tokenType] = fn
}

func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunction) {
	parser.infixParseFns[tokenType] = fn
}

// the engine function that parses the input program, constructs the
// complete Abstract Syntax Tree and returns the root [ast.Program]
// node.
func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.curToken.Type != token.EOF {
		statement := parser.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.nextToken()
	}

	return program
}

// parses statements as per the current token type and returns a
// [ast.Statement] node.
func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

// parses let statements and returns a [ast.LetStatement] node.
// supposed to be called when parser.curToken.Type == [token.LET].
func (parser *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := &ast.LetStatement{Token: parser.curToken}

	// if the let is not immediately followed by an identifier, return nil
	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	letStatement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	// if the identifier is not immediately followed by a =, return nil
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping the expressions until encounter a semicolon
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return letStatement
}

// helper for checking the curToken type is the expected type.
func (parser *Parser) curTokenIs(tok token.TokenType) bool {
	return parser.curToken.Type == tok
}

// helper for checking the next token type is the expected type.
func (parser *Parser) peekTokenIs(tok token.TokenType) bool {
	return parser.peekToken.Type == tok
}

// helper that moves parser to next token, if the next token has the
// expected token type. if successful, the function returns true,
// else false.
func (parser *Parser) expectPeek(tok token.TokenType) bool {
	if parser.peekTokenIs(tok) {
		parser.nextToken()
		return true
	} else {
		parser.peekError(tok)
		return false
	}
}

// parses return statements and returns a [ast.ReturnStatement] node.
// supposed to be called when parser.curToken.Type == [token.RETURN].
func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	returnStatement := &ast.ReturnStatement{Token: parser.curToken}

	parser.nextToken()

	// TODO: skipping the expressions until encounter a semicolon
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return returnStatement
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expStatement := &ast.ExpressionStatement{Token: parser.curToken}

	expStatement.Expression = parser.parseExpression(LOWEST)

	// optional semicolon so something like "5 + 5" can be typed in REPL
	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return expStatement
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefixParseFn := parser.prefixParseFns[parser.curToken.Type]
	if prefixParseFn == nil {
		return nil
	}

	leftExp := prefixParseFn()
	return leftExp
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
}
