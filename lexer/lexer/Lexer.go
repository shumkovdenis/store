package lexer

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/shumkovdenis/store/lexer/token"
)

type Lexer struct {
	Name   string
	Input  string
	Tokens chan token.Token
	State  LexFn
	Start  int
	Pos    int
	Width  int
}

func (lexer *Lexer) Emit(tokenType token.TokenType) {
	lexer.Tokens <- token.Token{Type: tokenType, Value: lexer.Input[lexer.Start:lexer.Pos]}
	lexer.Start = lexer.Pos
}

func (lexer *Lexer) Errorf(format string, args ...interface{}) LexFn {
	lexer.Tokens <- token.Token{
		Type:  token.TokenError,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

func (lexer *Lexer) InputToEnd() string {
	return lexer.Input[lexer.Pos:]
}

func (this *Lexer) IsEOF() bool {
	return this.Pos >= len(this.Input)
}

func (lexer *Lexer) Dec() {
	lexer.Pos--
}

func (lexer *Lexer) Inc() {
	lexer.Pos++
	if lexer.Pos >= utf8.RuneCountInString(lexer.Input) {
		lexer.Emit(token.TokenEOF)
	}
}

func (lexer *Lexer) Next() rune {
	if lexer.Pos >= utf8.RuneCountInString(lexer.Input) {
		lexer.Width = 0
		return token.EOF
	}

	result, width := utf8.DecodeRuneInString(lexer.Input[lexer.Pos:])

	lexer.Width = width
	lexer.Pos += lexer.Width
	return result
}

func (lexer *Lexer) SkipWhitespace() {
	for {
		ch := lexer.Next()

		if !unicode.IsSpace(ch) {
			lexer.Dec()
			break
		}

		if ch == token.EOF {
			lexer.Emit(token.TokenEOF)
			break
		}
	}
}
