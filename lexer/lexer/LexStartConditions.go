package lexer

import (
	"github.com/shumkovdenis/store/lexer/token"
)

func LexStartConditions(lexer *Lexer) LexFn {
	lexer.Pos += len(token.StartConditions)
	lexer.Emit(token.TokenStartConditions)
	return LexType
}
