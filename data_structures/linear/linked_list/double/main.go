package main

import "fmt"

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// perhatikan state linked list dan node
// keduanya harus sinkron

// menambahkan node baru setelah tail
func (dl *DoublyLinkedList) push(data interface{}) {
	node := &Node{data: data}
	if dl.head == nil {
		dl.head = node
		dl.tail = node
		return
	}

	// state new node
	node.prev = dl.tail

	// state tail
	dl.tail.next = node

	// state linked list
	dl.tail = node
}

func (dl *DoublyLinkedList) unshift(data interface{}) {
	// node := &Node{data: data, next: dl.head}
	node := &Node{}
	node.data = data
	// new node state
	node.next = dl.head

	// current head state
	dl.head.prev = node

	// new node state
	dl.head = node
}

func (dl *DoublyLinkedList) traverse() {
	fmt.Println("traverse")
	current := dl.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
}

func (dl *DoublyLinkedList) reverse() {
	fmt.Println("reverse")
	current := dl.tail
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.prev
	}
}

func main() {
	dl := &DoublyLinkedList{}
	dl.push(1)
	dl.push(2)
	dl.push(3)
	fmt.Println("----------------------------")
	dl.traverse()
	fmt.Println()
	dl.reverse()

}
