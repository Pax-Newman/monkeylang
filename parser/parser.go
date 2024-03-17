package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lex *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lex: l}

	// Preload cur and peek
	p.next()
	p.next()

	return p
}

func (p *Parser) next() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.Next()
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.next()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Set expression

	// NOTE: For now we skip expressions till a semicolon
	for !p.expectPeek(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) curIs(tok token.TokenType) bool {
	return p.curToken.Type == tok
}

func (p *Parser) peekIs(tok token.TokenType) bool {
	return p.peekToken.Type == tok
}

func (p *Parser) expectPeek(tok token.TokenType) bool {
	if p.peekIs(tok) {
		p.next()
		return true
	} else {
		return false
	}
}
