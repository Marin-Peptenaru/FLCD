package main

import (
	"fmt"
	"my-lang/parser"
)

func main() {
	st := parser.NewSymbolTable()
	index := st.SaveSymbol("var_1")

	fmt.Print(index)
	fmt.Println(st.GetSymbol(index))

	fmt.Println(st.SaveSymbol("var_1"))

}