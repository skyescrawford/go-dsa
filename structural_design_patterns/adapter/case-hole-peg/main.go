package main

import (
	"fmt"
	"math"
)

type Circle interface {
	getRadius() float64
}

type RoundHole struct {
	radius float64
}

func (r RoundHole) getRadius() float64 {
	return r.radius
}

func (r RoundHole) fits(c Circle) bool {
	return c.getRadius() <= r.getRadius()
}


type RoundPeg struct {
	radius float64
}

func (r RoundPeg) getRadius() float64 {
	return r.radius
}

type Square interface {
	getWidth() float64
}

type SquarePeg struct {
	width float64
}

func (s SquarePeg) getWidth() float64 {
	return s.width
}

type SquarePegAdapter struct {
	square Square
}

func (s SquarePegAdapter) getRadius() float64 {
	result := float64(s.square.getWidth() / math.Sqrt(2))
	fmt.Printf("%.2f\n", result)
	return result
}

func main() {
	roundHole := RoundHole{radius: 5}
	// roundPeg := RoundPeg{radius: 5}

	squarePeg := SquarePeg{width: 8}

	squareAdapter := SquarePegAdapter{square:squarePeg}

	fmt.Println(roundHole.fits(squareAdapter))

}