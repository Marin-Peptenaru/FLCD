package fa

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)
const stateREstr = "^[A-Z][A-Z0-9_]*$"
const symbolREstr = "^[a-z_0-9]+$"

// reads the contents of the file from the given filepath and tries to construct a finite automata from it
func FromFile(filepath string) (FiniteAutomata, error){
	fa := NewFA()

	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Could not read the contents of the input file: " + err.Error())
	}

	fileLines := strings.Split(string(fileContent), "\n")

	if len(fileLines) != 5 {
		log.Fatal("Invalid FA format in input file.\n")
	}
	states := parseStates(fileLines[0])

	for _, state := range states {
		fa.AddState(state)
	}

	symbols := parseAlphabet(fileLines[1])

	for _, symbol := range symbols {
		fa.AddSymbol(symbol)
	}

	initialState := parseStates(fileLines[3])[0]

	fa.SetInitialState(initialState)

	finalStates := parseStates(fileLines[4])

	for _, state := range finalStates {
		fa.AddFinalState(state)
	}

	parseTransitions(fa, fileLines[2])

	return fa, nil
}

func parseStates(str string) []State {
	strStates := strings.Split(str, ",")
	stateRE, err := regexp.Compile(stateREstr)

	if err != nil {
		log.Fatal("Could not compile regex for parsing states: " + err.Error())
	}

	states := make([]State, 0)

	for _, strState := range strStates {
		strState = strings.TrimSpace(strings.TrimRight(strState,"\n"))
		if  stateRE.MatchString(strState){
			states = append(states, State(strState))
		} else {
			log.Printf("Invalid state format: %s\n", strState)
		}
	}

	return states
}

func parseAlphabet(str string) []Symbol {
	strSymbols := strings.Split(str, ",")
	symbolRE, err := regexp.Compile(symbolREstr)

	if err != nil {
		log.Fatal("Could not compile regex for parsing symbols: " + err.Error())
	}

	symbols := make([]Symbol, 0)

	for _, strSymbol := range(strSymbols) {
		strSymbol = strings.TrimSpace(strings.TrimRight(strSymbol,"\n"))
		if  symbolRE.MatchString(strSymbol){
			symbols = append(symbols, Symbol(strSymbol))
		} else {
			log.Printf("Invalid symbol format: %s\n", strSymbol)
		}
	}

	return symbols
}


func parseTransitions(fa FiniteAutomata, str string){
	symbolRE, err := regexp.Compile(symbolREstr)

	if err != nil {
		log.Fatal("Could not compile regex for parsing symbols: " + err.Error())
	}

	stateRE, err := regexp.Compile(stateREstr)

	if err != nil {
		log.Fatal("Could not compile regex for parsing states: " + err.Error())
	}

	transitionsStr := strings.Split(str, ",")

	for _, transitionStr := range transitionsStr {
		state, symbol, nextState := parseTransisiton(transitionStr)

		if !stateRE.MatchString(state) || !symbolRE.MatchString(symbol) || !stateRE.MatchString(nextState){
			continue
		}

		fa.AddTransition(State(state), Symbol(symbol), State(nextState))
	}
}

func parseTransisiton(transitionStr string) (string, string, string) {

	split := strings.Split(transitionStr, ":")

	state := strings.TrimSpace(split[0])

	regexp := regexp.MustCompile("->")

	regexSplit := regexp.Split(split[1], -1)

	symbol, nextState := strings.TrimSpace(regexSplit[0]), strings.TrimSpace(regexSplit[1])

	return state, symbol, nextState
}