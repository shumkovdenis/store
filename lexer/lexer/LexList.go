package lexer

import (
	"strings"

	"github.com/shumkovdenis/store/lexer/errors"
	"github.com/shumkovdenis/store/lexer/token"
)

func LexList(lexer *Lexer) LexFn {
	if strings.HasPrefix(lexer.InputToEnd(), token.ListProperties) {
		lexer.Emit(token.TokenListProperties)
		return nil
	} else {
		return lexer.Errorf(errors.LexerErrorMissingProperties)
	}
}
