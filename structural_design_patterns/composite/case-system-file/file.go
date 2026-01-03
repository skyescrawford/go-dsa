package main

import "fmt"

type File struct {
	name string
}

// implement Component interface
func (f *File) search(name string) {
	fmt.Printf("searching in %s\n", f.name)
}
