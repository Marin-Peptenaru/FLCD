package parser

type LexicalVerifier interface {
	SyntacticallyCorrect(string) bool
	IsPredefinedToken(string) bool
	IsConstant(string) bool
	IsIndetifier(string) bool
}