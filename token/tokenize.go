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
	ASSIGN  = "="
	PLUS    = "+"
	ASTRISK = "*"
	SLASH   = "/"
	PERCENT = "%"
	GT      = ">"
	LT      = "<"
	BANG    = "!"
	EQ      = "=="
	NOTEQ   = "!="
	GE      = ">="
	LE      = "<="

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
	IF       = "IF"
	ELSE     = "ELSE"
	ELIF     = "ELIF"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var Keywords = map[string]TokenType{
	"function": FUNCTION,
	"sunbc":    LET,
	"oyebc":    START,
	"bhej":     RETURN,
	"bhagbc":   END,
	"jebc":     IF,
	"harjebc":  ELIF,
	"nistobc":  ELSE,
	"sahi":     TRUE,
	"galat":    FALSE,
}

func LookUpIdent(ident string) TokenType {
	keyword, exists := Keywords[ident]
	if exists {
		return keyword
	}
	return IDENT
}
