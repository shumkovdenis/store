package lexer

import (
	"strings"

	"github.com/shumkovdenis/store/lexer/errors"
	"github.com/shumkovdenis/store/lexer/token"
)

func LexBegin(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()

	if strings.HasPrefix(lexer.InputToEnd(), token.LeftParenthesis) {
		return LexLeftParenthesisConditions
	} else {
		return lexer.Errorf(errors.LexerErrorMissingLeftParenthesis)
	}
}
