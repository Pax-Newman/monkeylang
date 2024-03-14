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
	return nil
}
