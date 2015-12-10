package token

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF

	TokenLeftParenthesisConditions
	TokenRightParenthesisConditions

	TokenTypeString

	TokenLeftParenthesisArgument
	TokenRightParenthesisArgument

	TokenListProperties

	TokenLeftBracketKey
	TokenRightBracketKey

	TokenLeftDoubleQuotesKey
	TokenRightDoubleQuotesKey

	TokenEqualSign

	TokenLeftDoubleQuotesValue
	TokenRightDoubleQuotesValue

	TokenOperatorAnd

	TokenKey
	TokenValue
)
