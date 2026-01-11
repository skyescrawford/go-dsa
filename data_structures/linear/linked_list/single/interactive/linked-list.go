package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Push(data int) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
		return
	}

	current := l.head
	if current.next != nil {
		current = current.next
	}

	current.next = node
}

func (l *LinkedList) Unshift(data int) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
		return
	}

	node.next = l.head
	l.head = node
}

func (l *LinkedList) Iterate() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}

func (l *LinkedList) Reset() {
	l.head = nil
}
