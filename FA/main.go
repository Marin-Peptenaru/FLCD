package main

import (
	"finite-automata/fa"
	"fmt"
	"log"
	"strings"
)

func printMenu() {
	fmt.Printf("show -> show FA\ntest -> test if a string is accepted by the FA\nhelp -> show menu\nexit -> close program\n")
}

func test(automata fa.FiniteAutomata) {
	var str string
	fmt.Scanf("%s", &str)
	symbols := make([]fa.Symbol, 0)

	for _, char := range str {
		symbols = append(symbols, fa.Symbol(char))
	}

	accepted := automata.IsAccepted(symbols)

	if accepted {
		fmt.Printf("%s is accepted by the FA.\n", str)
	} else {
		fmt.Printf("%s is not accepted by the FA.\n", str)
	}
}
func main() {
	automata, err := fa.FromFile("FA.in")

	if err != nil {
		log.Fatal(err)
	}

	if !automata.IsDeterministic() {
		fmt.Println("The FA is not deterministic")
	} else {
		fmt.Println("The FA is deterministic")
	}

	printMenu()
	var command string
	for {
		fmt.Scanf("%s\n", &command)
		command := strings.ToLower(strings.TrimSpace(command))
		switch command {
		case "show":
			fmt.Println(automata)
		case "help":
			printMenu()
		case "exit":
			return
		case "test":
			test(automata)
		default:
			fmt.Println("Unknown command.")
		}
	}
}
