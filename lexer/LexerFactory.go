package lexer

import (
	"github.com/shumkovdenis/store/lexer/lexer"
	"github.com/shumkovdenis/store/lexer/token"
)

func BeginLexing(name, input string) *lexer.Lexer {
	l := &lexer.Lexer{
		Name:   name,
		Input:  input,
		State:  lexer.LexBegin,
		Tokens: make(chan token.Token, 3),
	}

	return l
}
