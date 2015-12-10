package token

import (
	"fmt"
)

type Token struct {
	Type  TokenType
	Value string
}

func (token Token) String() string {
	switch token.Type {
	case TokenEOF:
		return "EOF"
	case TokenError:
		return token.Value
	}

	return fmt.Sprintf("%q", token.Value)
}
