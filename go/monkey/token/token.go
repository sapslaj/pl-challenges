package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF                 = "EOF"
	IDENT               = "IDENT"
	INT                 = "INT"
	ASSIGN              = "="
	PLUS                = "+"
	COMMA               = ","
	SEMICOLON           = ";"
	LPAREN              = "("
	RPAREN              = ")"
	LBRACE              = "{"
	RBRACE              = "}"
	FUNCTION            = "FUNCTION"
	LET                 = "LET"
)

var Keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	tok, ok := Keywords[ident]
	if ok {
		return tok
	}
	return IDENT
}
