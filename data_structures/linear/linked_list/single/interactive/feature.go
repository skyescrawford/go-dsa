package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var inputPrompt = survey.Input{Message: "Data: "}

func addToEnd(list *LinkedList, data int) {
	survey.AskOne(&inputPrompt, &data)
	list.Push(data)
}

func addToStart(list *LinkedList, data int) {
	survey.AskOne(&inputPrompt, &data)
	list.Unshift(data)
}

func displayList(list *LinkedList) {
	if list.head == nil {
		fmt.Println("list is empty")
	} else {
		list.Iterate()
	}
	fmt.Print("\nPress enter to continue...")
	fmt.Scanln()
}

func clearList(list *LinkedList) {
	list.Reset()
}
