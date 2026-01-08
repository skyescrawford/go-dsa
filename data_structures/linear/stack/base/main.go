package main

import (
	"fmt"
	"os"
)

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{items: make([]interface{}, 0)}
}

func (s *Stack) Push(data interface{}) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() (result bool) {
	if len(s.items) == 0 {
		result = true
		return
	}
	result = false
	return
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Iterate() {
	fmt.Println("iterating... 🚀")
	for _, val := range s.items {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push("aaa")
	s.Push(true)
	s.Push(3.14)

	s.Iterate()

	data, success := s.Pop()
	if !success {
		fmt.Println("ayam 🐔")
		os.Exit(1)
	}

	fmt.Printf("data: %v\n", data)
	s.Iterate()

	fmt.Printf("peek: %v\n", s.Peek())

}
