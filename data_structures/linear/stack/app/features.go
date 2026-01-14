package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Add(s *Stack) {
	survey.AskOne(&addPrompt, &data)
	s.push(data)
}

func Peek(s *Stack) {
	node, ok := s.top()
	if !ok {
		fmt.Println("Stack is empty")
		fmt.Scanln()
		return
	}
	fmt.Printf("Peek: %d\n", node.value)
	fmt.Scanln()
}

func Take(s *Stack) {
	node, ok := s.pop()
	if !ok {
		fmt.Println("Stack is empty")
		fmt.Scanln()
		return
	}
	fmt.Printf("Popped: %d\n", node.value)
	fmt.Scanln()
}

func Size(s *Stack) {
	fmt.Printf("Size: %d\n", s.size)
	fmt.Scanln()
}

func Check(s *Stack) {
	fmt.Printf("Is empty? %t\n", s.isEmpty())
	fmt.Scanln()
}

func Reset(s *Stack) {
	if empty := s.isEmpty(); empty {
		fmt.Println("Stack is empty, nothing to clear out")
		fmt.Scanln()
		return
	}
	s.clear()
	fmt.Println("Stack is cleared")
	fmt.Scanln()
}

func Display(s *Stack) {
	s.display()
	fmt.Scanln()
}
