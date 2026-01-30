package ast

import (
	"github.com/iashyam/bclang/token"
)

type Node interface {
	TokenLitral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Satements []Statement
}

func (p *Program) TokenLitral() string {
	if len(p.Satements) > 0 {
		return p.Satements[0].TokenLitral()
	}
	return ""
}

func NewASTProgram() *Program {
	return &Program{}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

func (ls *LetStatement) TokenLitral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) StatementNode() {}

func (id *Identifier) TokenLitral() string {
	return id.Token.Literal
}

func (id *Identifier) StatementNode() {}
