package parser

import (
	"github.com/iashyam/bclang/ast"
	"github.com/iashyam/bclang/lexer"
	"github.com/iashyam/bclang/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{lexer: l}
	p.advanceTokens()
	p.advanceTokens()
	return &p
}

func (p *Parser) advanceTokens() {
	p.currentToken = p.nextToken
	p.nextToken = p.lexer.NextToken()
}

func (p *Parser) ParsePrograme() *ast.Program {
	return nil
}
