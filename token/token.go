package token

type TokenType string

const (
	ILLEGAL = TokenType("ILLEGAL")
	EOF     = TokenType("EOF")

	IDENT = TokenType("IDENT")
	INT   = TokenType("INT")

	// Operators
	ASSIGN = TokenType("=")
	PLUS   = TokenType("+")

	// Delimiters
	COMMA     = TokenType(",")
	SEMICOLON = TokenType(";")

	LPAREN = TokenType("(")
	RPAREN = TokenType(")")
	LBRACE = TokenType("{")
	RBRACE = TokenType("}")

	// Keywords
	FUNCTION = TokenType("FUNCTION")
	LET      = TokenType("LET")
)

var keyword = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keyword[ident]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Type TokenType

	Literal string
}
