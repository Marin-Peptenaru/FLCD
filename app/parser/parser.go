package parser

type Token string

type Parser interface {
	Parse(filepath string) (ProgramInternalForm, SymbolTable, SymbolTable)
}