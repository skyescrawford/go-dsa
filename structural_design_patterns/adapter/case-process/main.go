package main

import "fmt"

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (a Adapter) process() {
	fmt.Println("[Adapter] Connecting...")
	a.adaptee.convert()
	fmt.Println("[Adapter] Connected 🚀")
}

// adaptee adalah client
type Adaptee struct {
	adapterType int
}

func (a Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
