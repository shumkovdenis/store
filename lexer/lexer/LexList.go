package lexer

import (
	"strings"

	"github.com/shumkovdenis/store/lexer/errors"
	"github.com/shumkovdenis/store/lexer/token"
)

func LexList(lexer *Lexer) LexFn {
	if strings.HasPrefix(lexer.InputToEnd(), token.Properties) {
		return LexProperties
	} else {
		return lexer.Errorf(errors.LexerErrorMissingList)
	}
}
