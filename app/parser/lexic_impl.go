package parser

import (
	"log"
	"regexp"
)

type lexicImpl struct {
	separators       []string
	operators        []string
	predefinedTokens []string
	stringFormat     *regexp.Regexp
	charFormat       *regexp.Regexp
	intFormat        *regexp.Regexp
	identifierFormat *regexp.Regexp
}

func NewLexicalVerifier() LexicalVerifier {

	stringFormat, err := regexp.Compile(`^".*"$`)
	if err != nil {
		log.Fatal("Could not build regex for string checking: " + err.Error())
	}

	charFormat, err := regexp.Compile(`^'.'$`)
	if err != nil {
		log.Fatal("Could not build regex for char checking: " + err.Error())
	}

	intFormat, err := regexp.Compile(`^(0|[\+\-]?[1-9]\d*)$`)

	if err != nil {
		log.Fatal("Could not build regex for int checking: " + err.Error())
	}

	identfierFormat, err := regexp.Compile(`^[_a-zA-z][a-zA-Z0-9_]*$`)

	if err != nil {
		log.Fatal("Could not build regex for identifier checking: " + err.Error())
	}

	return &lexicImpl{
		separators:       []string{";", "(", ")", "[", "]", "{", "}", "#"},
		operators:        []string{"+", "-", "*", "/", ">", "<", "=", "<=", ">=", "!=", "==", "%"},
		predefinedTokens: []string{"if", "not", "and", "or", "else", "while", "for", "var", "begin", "end", "read", "print", "int", "char", "string"},
		stringFormat:     stringFormat,
		charFormat:       charFormat,
		intFormat:        intFormat,
		identifierFormat: identfierFormat,
	}
}

func (l *lexicImpl) IsOperator(token string) bool  { 
	for _, operator := range l.operators {
		if token == operator{
			return true;
		}
	}
	return false;
}

func (l *lexicImpl) IsSeparator(token string) bool { 
	for _, separator := range l.separators {
		if token == separator{
			return true;
		}
	}
	return false;
}

func (l *lexicImpl) IsPredefinedToken(token string) bool {
	for _, predefinedToken := range l.predefinedTokens {
		if token == predefinedToken {
			return true
		}
	}
	return false
}

func (l *lexicImpl) IsConstant(token string) bool {
	return l.IsChar(token) || l.IsString(token) || l.IsInt(token)
}

func (l *lexicImpl) IsIdentifier(token string) bool { return l.identifierFormat.MatchString(token) }

func (l *lexicImpl) IsChar(token string) bool       { return l.charFormat.MatchString(token) }

func (l *lexicImpl) IsInt(token string) bool        { return l.intFormat.MatchString(token) }

func (l *lexicImpl) IsString(token string) bool     { return l.stringFormat.MatchString(token) }
