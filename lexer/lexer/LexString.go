package lexer

import "github.com/shumkovdenis/store/lexer/token"

func LexString(lexer *Lexer) LexFn {
	lexer.Pos += len(token.String)
	lexer.Emit(token.TokenString)
	return LexList
}
