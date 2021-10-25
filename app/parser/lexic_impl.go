package parser

type lexicImpl struct {
}

func NewLexicalVerifier() LexicalVerifier {
	return &lexicImpl{}
}

func (l *lexicImpl) SyntacticallyCorrect(string) bool { return false }
func (l *lexicImpl) IsPredefinedToken(string) bool    { return false }
func (l *lexicImpl) IsConstant(string) bool           { return false }
func (l *lexicImpl) IsIndetifier(string) bool         { return false }