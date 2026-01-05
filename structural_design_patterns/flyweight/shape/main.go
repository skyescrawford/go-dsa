package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
	color string
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing %s circle\n", c.color)
}

type ShapeFactory struct {
	circles map[string]*Circle
}

func (s *ShapeFactory) GetCircle(color string) *Circle {
	if s.circles == nil {
		s.circles = make(map[string]*Circle)
	}

	if s.circles[color] == nil {
		s.circles[color] = &Circle{color: color}
	}

	return s.circles[color]
}

func main() {
	shapeFactory := &ShapeFactory{}

	colors := []string{"red", "green", "blue", "purple", "yellow", "grey", "gold", "silver"}

	for _, color := range colors {
		circle := shapeFactory.GetCircle(color)
		circle.Draw()
	}

	fmt.Println("------------------------------------------------------")

	redCircle := shapeFactory.GetCircle("red")
	redCircle.Draw()

	goldCircle := shapeFactory.GetCircle("gold")
	goldCircle.Draw()
}
