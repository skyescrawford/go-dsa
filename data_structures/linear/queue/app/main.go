package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
)

var commandPrompt = survey.Select{Message: "Select command:", Options: []string{"join queue", "display", "who is next?", "proceed queue", "is empty?", "size", "clear", "exit"}}
var command string

func main() {
	fmt.Println("Welcome to interactive queue")
	queue := NewQueue()

	for {
		clearScreen()

		survey.AskOne(&commandPrompt, &command)

		switch command {
		case "join queue":
			JoinToQueue(queue)
		case "display":
			Display(queue)
		case "who is next?":
			PeekWhosNext(queue)
		case "proceed queue":
			ProceedQueue(queue)
		case "is empty?":
			IsEmpty(queue)
		case "size":
			Size(queue)
		case "clear":
			Reset(queue)
		case "exit":
			os.Exit(1)
		default:
			fmt.Println("Invalid command")
			os.Exit(1)
		}
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
