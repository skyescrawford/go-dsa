package main

import "fmt"

// abstraction
type DrawAPI interface {
	DrawCircle(radius, x, y int)
}

type RedCircle struct{}

// implementation
func (rc *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing red circle of radius %d at (%d, %d)\n", radius, x, y)
}

// abstraction
type Shape interface {
	Draw()
}

type Circle struct {
	x, y, radius int
	drawApi      DrawAPI
}

func (c *Circle) Draw() {
	c.drawApi.DrawCircle(c.radius, c.x, c.y)
}

// Client
func draw(d Shape) {
	d.Draw()
}

func main() {
	redCircle := &Circle{100, 100, 10, &RedCircle{}}
	draw(redCircle)
}
