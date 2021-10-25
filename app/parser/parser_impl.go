package parser

type parserImpl struct {
	lexic           LexicalVerifier
	constantTable   SymbolTable
	identifierTable SymbolTable
}