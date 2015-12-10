package lexer

import (
	"strings"

	"github.com/shumkovdenis/store/lexer/errors"
	"github.com/shumkovdenis/store/lexer/token"
)

func LexProperties(lexer *Lexer) LexFn {
	lexer.Pos += len(token.Properties)
	lexer.Emit(token.TokenProperties)

	if strings.HasPrefix(lexer.InputToEnd(), token.StartKey) {
		return LexStartKey
	} else {
		return lexer.Errorf(errors.LexerErrorMissingKey)
	}
}
