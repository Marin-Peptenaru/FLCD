package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type parserImpl struct {
	lexic LexicalVerifier
}

func NewParser() Parser {
	return &parserImpl{lexic: NewLexicalVerifier()}
}

func separateTokens(s string) []string {
	tokens := make([]string, 0)

	separators, err := regexp.Compile(`[\+\-\*/%!(!=)\s\t==\n(<=)(>=)=<>#\(\)\{\}\[\];]`)

	if err != nil {
		log.Fatal("Could not create separators regex for detecting tokens: " + err.Error())
		return tokens
	}

	separatorMatches := separators.FindAllIndex([]byte(s), -1)

	for index, separatorIndexes := range separatorMatches {
		separatorStart := separatorIndexes[0]
		separatorEnd := separatorIndexes[1]

		// the first separator should be a # and should be at the start of the program,
		// so for the first separator I do not have a token before it
		if index == 0 {
			tokens = append(tokens, s[separatorStart:separatorEnd])
			continue
		}

		// for any other sepator,
		// I add it to the token list but before it I add the token that was between the current and the previous separator
		previousSeparatorEnd := separatorMatches[index-1][1]

		previousToken := s[previousSeparatorEnd:separatorStart]
		separator := s[separatorStart:separatorEnd]

		if len(previousToken) > 0 {
			tokens = append(tokens, previousToken)
		}

		if len(strings.TrimSpace(separator)) > 0 {
			tokens = append(tokens, separator)
		}

	}
	fmt.Println("Tokens: ", tokens)
	return tokens
}

func shouldLookAhead(token string) bool {
	return token == "!" || token == "<" || token == ">" || token == "="
}

func (p *parserImpl) Parse(filepath string) (ProgramInternalForm, SymbolTable, SymbolTable) {
	pif := NewPIF()
	constants := NewSymbolTable()
	identifiers := NewSymbolTable()
	lexicalErrors := strings.Builder{}

	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal("could not read program file: ", err)
	}

	tokens := separateTokens(string(file))
	log.Printf("Separated tokens: %v\n", tokens)

	programLenght := len(tokens)

	if tokens[0] != "#" || tokens[programLenght-1] != "#" {
		log.Fatal("Missing program delimiters")
	}

	for i := 0; i < programLenght; i++ {
		token := tokens[i]

		// check for operators such as >=, <=, != , ==
		if shouldLookAhead(token) && i < programLenght-1 {
			if tokens[i+1] == "=" {
				token += tokens[i+1]
				i++
			}
		}

		// check for strings and chars
		if token == "'" || token == "\"" {
			for i = i + 1; i < programLenght; i++ {
				token += tokens[i]

				// if I have found another ' or "
				if tokens[i] == string(token[0]) {
					break
				}
			}
		}

		if p.lexic.IsPredefinedToken(token) || p.lexic.IsOperator(token) || p.lexic.IsSeparator(token) {
			pif.Add(token, NullIndex())
		} else if p.lexic.IsConstant(token) {
			index := constants.SaveSymbol(Symbol(token))
			pif.Add(token, index)
		} else if p.lexic.IsIdentifier(token) {
			index := identifiers.SaveSymbol(Symbol(token))
			pif.Add(token, index)
		} else {
			lexicalErrors.WriteString(fmt.Sprintf("Lexically invalid token: %s\n", token))
		}
	}
	if lexicalErrors.Len() > 0 {
		log.Fatal(lexicalErrors.String())
	}

	return pif, constants, identifiers
}
