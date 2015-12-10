package lexer

import (
	"github.com/shumkovdenis/store/lexer/token"
)

func LexLeftParenthesisConditions(lexer *Lexer) LexFn {
	lexer.Pos += len(token.LeftParenthesis)
	lexer.Emit(token.TokenLeftParenthesisConditions)
	return LexType
}
