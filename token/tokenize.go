package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal byte) Token {
	return Token{Type: tokenType, Literal: string(literal)}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//IDENTIFIERS and LITERALS
	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	START    = "START"
	END      = "END"
	RETURN   = "RETURN"
)

var Keywords = map[string]TokenType{
	"function": FUNCTION,
	"sunbc":    LET,
	"oyebc":    START,
	"bhej":     RETURN,
	"bhagbc":   END,
}

func LookUpIdent(ident string) TokenType {
	keyword, exists := Keywords[ident]
	if exists {
		return keyword
	}
	return IDENT
}
