package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
)

var commandPrompt = survey.Select{Message: "Select option: ", Options: []string{"push", "top", "pop", "display", "is empty?", "size", "clear", "exit"}}
var command string

var addPrompt = survey.Input{Message: "Enter data:"}
var data int

func main() {
	stack := NewStack()
	for {
		clearScreen()
		survey.AskOne(&commandPrompt, &command)

		switch command {
		case "push":
			Add(stack)
		case "top":
			Peek(stack)
		case "pop":
			Take(stack)
		case "display":
			Display(stack)
		case "is empty?":
			Check(stack)
		case "size":
			Size(stack)
		case "clear":
			Reset(stack)
		case "exit":
			os.Exit(1)
		default:
			fmt.Println("Invalid option")
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
_:
	cmd.Run()
}
