package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Add(set *Set) {
	var data string
	addCommand := survey.Input{Message: "Enter data to add: "}
	survey.AskOne(&addCommand, &data)
	ok := set.Add(data)
	if !ok {
		fmt.Println("Item already exist")
		fmt.Println("Press enter to continue...")
		fmt.Scanln()
		return
	}
	fmt.Println("Added successfully")
	fmt.Println("Press enter to continue...")
	fmt.Scanln()
}

func Delete(set *Set) {
	var data string
	deleteCommand := survey.Input{Message: "Enter data to delete: "}
	survey.AskOne(&deleteCommand, &data)
	ok := set.Delete(data)
	if !ok {
		fmt.Printf("Item %s is not exist in set\n", data)
		fmt.Println("Press enter to continue...")
		fmt.Scanln()
		return
	}
	fmt.Println("Deleted successfully")
	fmt.Print("\nPress enter to continue...")
	fmt.Scanln()
}

func Display(set *Set) {
	set.Display()
	fmt.Print("\nPress enter to continue...")
	fmt.Scanln()
}

func Size(set *Set) {
	fmt.Printf("%d\n", set.Size())
	fmt.Print("\nPress enter to continue...")
	fmt.Scanln()
}

func IsEmpty(set *Set) {
	fmt.Printf("Is empty? %t\n", set.IsEmpty())
	fmt.Print("\nPress enter to continue...")
	fmt.Scanln()
}

func Reset(set *Set) {
	set.Clear()
	fmt.Println("Set has cleared")
	fmt.Print("\nPress enter to continue...")
	fmt.Scanln()
}
