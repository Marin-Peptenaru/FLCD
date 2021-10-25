package parser

type Parser interface {
	Parse(filepath string) (ProgramInternalForm, SymbolTable, SymbolTable, string)
}