package parser

type LexicalVerifier interface {
	IsOperator(string) bool
	IsSeparator(string) bool
	IsPredefinedToken(string) bool
	IsConstant(string) bool
	IsIdentifier(string) bool
	IsChar(string) bool
	IsInt(string) bool
	IsString(string) bool
}