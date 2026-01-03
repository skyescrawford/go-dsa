package main

func main() {
	red := &Red{}
	blue := &Blue{}
	green := &Green{}
	purple := &Purple{}

	redCircle := &Circle{color: red}
	blueSquare := &Square{color: blue}
	greenCircle := &Circle{color: green}
	redTriangle := &Triangle{color: red}
	purpleRectangle := &Rectangle{color: purple}

	purpleRectangle.draw()
	redCircle.draw()
	blueSquare.draw()
	greenCircle.draw()
	redTriangle.draw()
}
