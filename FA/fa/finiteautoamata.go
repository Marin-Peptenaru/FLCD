package fa

import (
	"fmt"
)

type State string
type Symbol string

type Configuration struct {
	state  State
	symbol Symbol
}

type FiniteAutomata interface {
	NextState(Configuration) (State, error)
	IsAccepted([]Symbol) bool
	IsDeterministic() bool
	AddState(State)
	AddSymbol(Symbol)
	SetInitialState(State)
	AddFinalState(State)
	AddTransition(State, Symbol, State)
}

type faImpl struct {
	initialState State
	states       map[State]bool
	alphabet     map[Symbol]bool
	finalStates  map[State]bool
	transitions  map[Configuration][]State
}

func NewFA() FiniteAutomata {
	return &faImpl{
		initialState: State(""),
		states:       make(map[State]bool),
		alphabet:     make(map[Symbol]bool),
		finalStates:  make(map[State]bool),
		transitions:  make(map[Configuration][]State),
	}
}

func (fa *faImpl) NextState(c Configuration) (State, error) {
	states, ok := fa.transitions[c]

	_, stateIsValid := fa.states[c.state]
	_, symbolIsValid := fa.alphabet[c.symbol]

	if !stateIsValid || !symbolIsValid {
		return c.state, fmt.Errorf("invalid configuration, either state is not accepted or symbol does not belong to the alphabet %v", c)
	}

	if !ok || len(states) > 1 {
		return c.state, fmt.Errorf("could not compute next state for configuration %v", c)
	}
	return states[0], nil
}

func (fa *faImpl) IsAccepted(sequence []Symbol) bool {
	configuration := Configuration{}
	state := fa.initialState
	err := error(nil)

	for _, symbol := range sequence {
		configuration.state = state
		configuration.symbol = symbol

		state, err = fa.NextState(configuration)

		if err != nil {
			return false
		}
	}
	_, stateIsFinal := fa.finalStates[state]

	return stateIsFinal
}

func (fa *faImpl) IsDeterministic() bool {

	for _, states := range fa.transitions {
		if len(states) > 1 {
			return false
		}
	}
	return true
}

func (fa *faImpl) AddState(newState State) {
	fa.states[newState] = true
}
func (fa *faImpl) AddSymbol(newSymbol Symbol)      {
	fa.alphabet[newSymbol] = true
}
func (fa *faImpl) SetInitialState(newInitialState State) {
	fa.states[newInitialState] = true
	fa.initialState = newInitialState
}
func (fa *faImpl) AddFinalState(finalState State)   {
	fa.states[finalState] = true
	fa.finalStates[finalState] = true
}

func (fa *faImpl) AddTransition(state State, symbol Symbol, nextState State) {
	conf := Configuration{state: state, symbol: symbol}

	states, hasAListOfTransitions := fa.transitions[conf]

	if !hasAListOfTransitions {
		fa.transitions[conf] = []State{nextState}
		return
	}

	for _, s := range states {
		if s == nextState{
			return
		}
	}

	fa.transitions[conf] = append(states, nextState)
}

func (fa *faImpl) String() string {
	initialState := string(fa.initialState)

	states := ""

	for state, _ := range fa.states {
		states += fmt.Sprintf(" %s", state)
	}

	alphabet := " "

	for symbol, _ := range fa.alphabet {
		alphabet += fmt.Sprintf(" %s", symbol)
	}

	finalStates := " "

	for state, _ := range fa.finalStates {
		finalStates += fmt.Sprintf(" %s", state)
	}

	transitions := ""
	
	for conf, states := range fa.transitions {
		transitions += fmt.Sprintf("(%v,%v)->", conf.state, conf.symbol)

		for _, state := range states {
			transitions += fmt.Sprintf("%v|", state)
		}

		transitions += "\n"
	}

	return fmt.Sprintf(
		"Init. State: %s\nStates: %s\nAlphabet: %s\nFinal States: %s\nTranisitions:\n%s",
		initialState, states, alphabet, finalStates, transitions)
	
} 



	