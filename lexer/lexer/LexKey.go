package lexer

import (
	"strings"

	"github.com/shumkovdenis/store/lexer/errors"
	"github.com/shumkovdenis/store/lexer/token"
)

func LexKey(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LexerErrorMissingKey)
		}

		if strings.HasPrefix(lexer.InputToEnd(), token.EndKey) {
			lexer.Emit(token.TokenKey)
			return LexEndKey
		}

		lexer.Inc()
	}
}
