package lexer

import "github.com/shumkovdenis/store/lexer/token"

func LexEndKey(lexer *Lexer) LexFn {
	lexer.Pos += len(token.EndKey)
	lexer.Emit(token.TokenEndKey)
	return nil
}
