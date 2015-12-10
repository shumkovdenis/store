package token

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF

	TokenStartConditions
	TokenEndConditions

	TokenString

	TokenProperties

	TokenStartKey
	TokenEndKey
	TokenKey

	TokenStartValue
	TokenEndValue
	TokenValue

	TokenEqualSign

	TokenAnd
)
