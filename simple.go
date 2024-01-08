package main

import "fmt"

func simpleSolution() {
	fmt.Println("Simple Condition")

	s1 := NewState("s1", false)
	s2 := NewState("s2", false)
	s3 := NewState("s3", true)

	s1.AddTransition('a', s2)
	s2.AddTransition('b', s3)
	s3.AddTransition('b', s2)
	s3.AddTransition('a', s3)

	automaton := NewAutomaton(s1)
	automaton.AddState(s2)
	automaton.AddState(s3)

	fmt.Println("ab", automaton.Accepts("ab"))
	fmt.Println("abb", automaton.Accepts("abb"))
	fmt.Println("aba", automaton.Accepts("aba"))
}
