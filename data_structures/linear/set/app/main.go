package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	set := New()

	for {
		clearScreen()

		var command string
		commandPrompt := survey.Select{Message: "Select command:", Options: []string{"add", "delete", "display", "size", "is empty?", "clear", "exit"}}
		survey.AskOne(&commandPrompt, &command)

		switch command {
		case "add":
			Add(set)
		case "delete":
			Delete(set)
		case "display":
			Display(set)
		case "size":
			Size(set)
		case "is empty?":
			IsEmpty(set)
		case "clear":
			Reset(set)
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
