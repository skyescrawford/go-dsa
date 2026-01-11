package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	list := NewList()
	commandPrompt := survey.Select{Message: "Select command:", Options: []string{"add to end", "add to start", "display list", "clear list", "exit"}}
	var command string

	for {
		clearScreen()
		survey.AskOne(&commandPrompt, &command)
		var data int
		switch command {
		case "add to end":
			addToEnd(list, data)
		case "add to start":
			addToStart(list, data)
		case "display list":
			displayList(list)
		case "clear list":
			clearList(list)
		case "exit":
			os.Exit(1)
		default:
			fmt.Println("invalid command")
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
