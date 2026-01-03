package main

import "fmt"

type Color interface {
	fill(shape string)
}

func drawLine(length int) {
	line := ""
	for i := 0; i < length; i++ {
		line += "-"
	}
	fmt.Printf("%s\n\n", line)

}

type Red struct{}

func (r *Red) fill(shape string) {
	fmt.Printf("fill %s with red\n", shape)
	drawLine(20)
}

type Blue struct{}

func (b *Blue) fill(shape string) {
	fmt.Printf("fill %s with blue\n", shape)
	drawLine(20)
}

type Green struct{}

func (g *Green) fill(shape string) {
	fmt.Printf("fill %s with green\n", shape)
	drawLine(20)
}

type Purple struct{}

func (p *Purple) fill(shape string) {
	fmt.Printf("fill %s with purple\n", shape)
	drawLine(20)
}
