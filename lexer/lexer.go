package lexer

import (
	"fmt"

	"github.com/dassongh/custom-interpreter/token"
)

// TODO: add support for Unicode characters
type Lexer struct {
	input        string // text to be tokenized
	position     int    // pointer to the current char
	readPosition int    // pointer to the next char
	character    byte   // represents ASCII character
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lexer.character {
	case '=':
		tok = newToken(token.ASSIGN, lexer.character)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.character)
	case '(':
		tok = newToken(token.LPAREN, lexer.character)
	case ')':
		tok = newToken(token.RPAREN, lexer.character)
	case ',':
		tok = newToken(token.COMMA, lexer.character)
	case '{':
		tok = newToken(token.LBRACE, lexer.character)
	case '}':
		tok = newToken(token.RBRACE, lexer.character)
	case '+':
		tok = newToken(token.PLUS, lexer.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.character) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		// ASCII code for the "NUL" character
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.character) {
		lexer.readChar()
	}

	result := lexer.input[position:lexer.position]

	fmt.Println("Result", result)

	return result
}

func isLetter(char byte) bool {
	// comparing ASCII numeric values
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
