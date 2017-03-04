package token

// TokenType allows us to distinguish between different types of tokens.
type TokenType string

// Token express each token.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL signifies a token/character we don't know about.
	ILLEGAL = "ILLEGAL"
	// EOF signifies end of file.
	EOF = "EOF"

	// IDENT signifies identifiers like foobar, x, y
	IDENT = "IDENT"
	// INT signifies integers like 1, 2, 34
	INT = "INT"

	// ASSIGN assigns a value to variables.
	ASSIGN = "="
	// PLUS signifies plus.
	PLUS = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
