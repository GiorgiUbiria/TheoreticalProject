package main

import "fmt"

type State struct {
	name        string
	transitions map[rune]*State
	isFinal     bool
}

func NewState(name string, isFinal bool) *State {
	return &State{name, make(map[rune]*State), isFinal}
}

func (s *State) AddTransition(symbol rune, state *State) {
	s.transitions[symbol] = state
}

type Automaton struct {
	states []*State
	start  *State
}

func NewAutomaton(start *State) *Automaton {
	return &Automaton{[]*State{start}, start}
}

func (a *Automaton) AddState(state *State) {
	a.states = append(a.states, state)
}

func (a *Automaton) Accepts(input string) bool {
	current := a.start
	for _, symbol := range input {
		next, ok := current.transitions[symbol]
		if !ok {
			fmt.Println("incorrect symbol")
			return false
		}
		current = next
	}

	return current.isFinal
}
