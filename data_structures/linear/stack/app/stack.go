package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(data int) {
	node := &Node{value: data}
	s.size++

	if s.head == nil {
		s.head = node
		return
	}

	next := s.head
	node.next = next
	s.head = node

}
func (s *Stack) top() (*Node, bool) {
	if s.head == nil {
		return nil, false
	}
	return s.head, true
}
func (s *Stack) pop() (*Node, bool) {
	if s.head == nil {
		return nil, false
	}
	head := s.head
	s.head = head.next
	s.size--
	return head, true
}
func (s *Stack) isEmpty() bool {
	return s.head == nil
}

func (s *Stack) clear() {
	s.head = nil
	s.size = 0
}

func (s *Stack) display() {
	if s.head == nil {
		fmt.Println("stack is empty")
		return
	}
	current := s.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
}
