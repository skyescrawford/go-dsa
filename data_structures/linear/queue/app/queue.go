package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type Queue struct {
	head *Node
	size int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) enqueue(value string) {
	node := &Node{value: value}
	q.size++
	if q.head == nil {
		q.head = node
		return
	}

	current := q.head
	for current.next != nil {
		current = current.next
	}

	current.next = node
}

func (q *Queue) peek() (*Node, bool) {
	if q.head == nil {
		return nil, false
	}

	return q.head, true
}

func (q *Queue) dequeue() (*Node, bool) {
	if q.head == nil {
		return nil, false
	}

	q.size--
	node := q.head
	q.head = q.head.next
	return node, true
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) display() string {
	var result string
	current := q.head
	for current != nil {
		result += fmt.Sprintf("%s ", current.value)
		current = current.next
	}
	return result
}

func (q *Queue) clear() {

	if q.head == nil {
		fmt.Println("Queue is empty, there is nothing to clear")
		return
	}
	q.head = nil
	q.size = 0
	fmt.Println("Queue is cleared")
}
