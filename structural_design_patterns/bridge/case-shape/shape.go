package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct {
	color Color
}

func (c *Circle) draw() {
	shape := "circle"
	fmt.Printf("draw %s\n", shape)
	c.color.fill(shape)
}

type Square struct {
	color Color
}

func (s *Square) draw() {
	shape := "square"
	fmt.Printf("draw %s\n", shape)
	s.color.fill(shape)
}

type Triangle struct {
	color Color
}

func (t *Triangle) draw() {
	shape := "triangle"
	fmt.Printf("draw %s\n", shape)
	t.color.fill(shape)
}

type Rectangle struct {
	color Color
}

func (r *Rectangle) draw() {
	shape := "rectangle"
	fmt.Printf("draw %s\n", shape)
	r.color.fill(shape)
}
