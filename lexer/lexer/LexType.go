package lexer

import (
	"strings"

	"github.com/shumkovdenis/store/lexer/errors"
	"github.com/shumkovdenis/store/lexer/token"
)

func LexType(lexer *Lexer) LexFn {
	if strings.HasPrefix(lexer.InputToEnd(), token.String) {
		return LexString
	} else {
		return lexer.Errorf(errors.LexerErrorMissingType)
	}
}
