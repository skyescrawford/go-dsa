package main

import "fmt"

type Folder struct {
	name       string
	components []Component
}

// implement Component interface
func (f *Folder) search(name string) {
	fmt.Printf("searching %s in %s\n", name, f.name)
	for _, composite := range f.components {
		composite.search(name)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
