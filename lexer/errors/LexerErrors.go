package errors

const (
	// LEXER_ERROR_UNEXPECTED_EOF        string = "Unexpected end of file"
	// LEXER_ERROR_MISSING_RIGHT_BRACKET string = "Missing a closing section bracket"
	LexerErrorMissingLeftParenthesis  string = "Missing a opening parenthesis"
	LexerErrorMissingRightParenthesis string = "Missing a closing parenthesis"
	LexerErrorMissingType             string = "Missing type"
	LexerErrorMissingProperties       string = "Missing properties"
)
