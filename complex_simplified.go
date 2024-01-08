package main

import (
	"fmt"
)

func complexSimplifiedSolution() {
	fmt.Println("Simplified Complex Solution")

	s1 := NewState("s1", false)
	s2 := NewState("s2", false)
	s3 := NewState("s3", false)
	s4 := NewState("s4", false)
	s5 := NewState("s5", true)

	s1.AddTransition('a', s2)
	s1.AddTransition('b', s1)
	s2.AddTransition('a', s3)
	s2.AddTransition('b', s1)
	s3.AddTransition('a', s3)
	s3.AddTransition('b', s4)
	s4.AddTransition('a', s2)
	s4.AddTransition('b', s5)
	s5.AddTransition('a', s5)
	s5.AddTransition('b', s5)

	automaton := NewAutomaton(s1)
	automaton.AddState(s2)
	automaton.AddState(s3)
	automaton.AddState(s4)
	automaton.AddState(s5)

	fmt.Println("ab", automaton.Accepts("ab"))
	fmt.Println("aab", automaton.Accepts("aab"))
	fmt.Println("aabb", automaton.Accepts("aabb"))
	fmt.Println("aabba", automaton.Accepts("aabba"))
}
