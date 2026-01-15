package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var addCommand = survey.Input{Message: "Enter data:"}
var data string

func JoinToQueue(q *Queue) {
	survey.AskOne(&addCommand, &data)
	q.enqueue(data)
}

func PeekWhosNext(q *Queue) {
	node, ok := q.peek()
	if !ok {
		fmt.Println("Queue is empty")
		fmt.Scanln()
		return
	}
	fmt.Printf("Next: %s\n", node.value)
	fmt.Scanln()
}

func ProceedQueue(q *Queue) {
	node, ok := q.dequeue()
	if !ok {
		fmt.Println("Queue is empty")
		fmt.Scanln()
		return
	}
	fmt.Printf("Proceeding: %s\n", node.value)
	fmt.Scanln()
}

func IsEmpty(q *Queue) {
	fmt.Printf("Is empty? %t\n", q.isEmpty())
	fmt.Scanln()
}

func Size(q *Queue) {
	fmt.Printf("Size: %d\n", q.size)
	fmt.Scanln()
}

func Display(q *Queue) {
	if q.size == 0 {
		fmt.Println("Queue is empty, there is nothing to show")
		fmt.Scanln()
		return
	}
	fmt.Printf("Display: %s\n", q.display())
	fmt.Scanln()
}

func Reset(q *Queue) {
	q.clear()
	fmt.Scanln()
}
