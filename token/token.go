package token

type Token struct {
	Type  TokenType
	Value string
}

//go:generate stringer -type=TokenType
type TokenType uint8

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifers & literals
	IDENT // [a-zA-Z]+
	INT   // [0-9]+

	// Operators
	ASSIGN   // =
	PLUS     // +
	MINUS    // -
	BANG     // !
	ASTERISK // *
	SLASH    // /
	EQ       // ==
	NOT_EQ   // !=

	LT // <
	GT // >

	// Delimiters
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	IF
	ELSE
	TRUE
	FALSE
	RETURN
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func New(tokType TokenType, ch byte) Token {
	return Token{Type: tokType, Value: string(ch)}
}

func Lookup(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
