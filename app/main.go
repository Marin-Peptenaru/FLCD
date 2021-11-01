package main

import (
	"fmt"
	"my-lang/parser"
)

func main() {
	parser := parser.NewParser()
	pif, constants, indetifiers := parser.Parse("p4.myl")
	fmt.Println("PIF")
	fmt.Println(pif)
	fmt.Println("Constants")
	fmt.Println(constants)
	fmt.Println("Identifiers")
	fmt.Println(indetifiers)

}