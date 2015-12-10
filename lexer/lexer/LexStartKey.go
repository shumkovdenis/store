package lexer

import "github.com/shumkovdenis/store/lexer/token"

func LexStartKey(lexer *Lexer) LexFn {
	lexer.Pos += len(token.StartKey)
	lexer.Emit(token.TokenStartKey)
	return LexKey
}
