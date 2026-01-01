package main

import "fmt"

// abstraction
type IDrawShape interface {
	drawShape(x, y [5]float32)
}

type ICountour interface {
	drawContour(x, y [5]float32)
	resizeByFactor(factor int)
}

// implementation
type DrawShape struct{}

func (d DrawShape) drawShape(x, y [5]float32) {
	fmt.Println("Drawing Shape")
}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (c DrawContour) drawContour(x, y [5]float32) {
	fmt.Println("Drawing Contour")
	c.shape.drawShape(c.x, c.y)
}

func (c DrawContour) resizeByFactor(factor int) {
	c.factor = factor
}

func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var countour ICountour = DrawContour{x, y, DrawShape{}, 2}
	countour.drawContour(x, y)
	// countour.resizeByFactor(2)
}
