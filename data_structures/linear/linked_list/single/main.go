package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) addToTail(data interface{}) {
	node := &Node{data: data}

	// case kalo empty list
	if l.head == nil {
		l.head = node // initiate first element of the list
		return
	}

	current := l.head
	// kalo ga empty ya looping lah
	for current.next != nil {
		current = current.next
	}

	current.next = node
}

func (l *LinkedList) cutTail() {
	secondLastItem := l.secondLastItem()
	secondLastItem.next = nil
}

func (l *LinkedList) lastItem() *Node {
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current
}

// punya second last item jika panjang list >= 2
func (l *LinkedList) secondLastItem() *Node {
	if l.head == nil || l.head.next == nil {
		return nil
	}

	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	return current
}

func (l *LinkedList) addToHead(data interface{}) {
	node := &Node{data: data, next: l.head}

	l.head = node
}

func (l *LinkedList) cutHead() {
	l.head = l.head.next
}

func (l *LinkedList) traverse() {
	current := l.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
}

func (l *LinkedList) reverse() {
	var prev *Node

	fmt.Println("reversing...")

	current := l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev

}

func main() {
	ll := &LinkedList{}
	ll.addToTail(1)
	ll.addToTail(2)
	ll.addToTail(3)

	ll.reverse()

	ll.traverse()
}
