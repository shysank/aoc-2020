package calc

import (
	"fmt"
	"strconv"
)

type Lexer struct {
	pos  int
	expr string
}

func NewLexer(e string) Lexer {
	return Lexer{expr: e}
}

func (lexer *Lexer) NextToken() (Token, error) {
	if lexer.reachedEnd() {
		return Eof(), nil
	}
	if " " == lexer.currentChar() {
		lexer.ignoreWhitespace()
	}
	if "+" == lexer.currentChar() {
		lexer.advance()
		return Plus(), nil
	}
	if "-" == lexer.currentChar() {
		lexer.advance()
		return Minus(), nil
	}
	if "*" == lexer.currentChar() {
		lexer.advance()
		return Mul(), nil
	}
	if "/" == lexer.currentChar() {
		lexer.advance()
		return Div(), nil
	}
	if "(" == lexer.currentChar() {
		lexer.advance()
		return LParen(), nil
	}
	if ")" == lexer.currentChar() {
		lexer.advance()
		return RParen(), nil
	}
	if isDigit(lexer.currentChar()) {
		return lexer.getIntegerToken()
	}
	return Token{}, fmt.Errorf("Lexer Error: Unparseable character %s", lexer.currentChar())
}

func (lexer *Lexer) currentChar() string {
	return lexer.expr[lexer.pos : lexer.pos+1]
}

func (lexer *Lexer) reachedEnd() bool {
	return lexer.pos == len(lexer.expr)
}

func (lexer *Lexer) advance() {
	lexer.pos++
}

func (lexer *Lexer) ignoreWhitespace() {
	for lexer.currentChar() == " " {
		lexer.advance()
	}
}

func (lexer *Lexer) getIntegerToken() (Token, error) {
	num := ""
	for !lexer.reachedEnd() && isDigit(lexer.currentChar()) {
		num = num + lexer.currentChar()
		lexer.advance()
	}

	integerVal, _ := strconv.Atoi(num)
	return NewToken(INTEGER, integerVal), nil
}

func isDigit(char string) bool {
	if _, err := strconv.Atoi(char); err != nil {
		return false
	}
	return true
}
