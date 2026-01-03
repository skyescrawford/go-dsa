package main

type Component interface {
	Operation()
}

func main() {
	l1 := &Leaf{name: "leaf 1"}
	l2 := &Leaf{name: "leaf 2"}
	l3 := &Leaf{name: "leaf 3"}
	l4 := &Leaf{name: "leaf 4"}
	l5 := &Leaf{name: "leaf 5"}
	l6 := &Leaf{name: "leaf 6"}
	l7 := &Leaf{name: "leaf 7"}

	c1 := &Composite{}
	c1.add(l3, l2, l1)

	c2 := &Composite{}
	c2.add(l4, l5)

	root := &Composite{}
	root.add(l6, l7, c1, c2)

	root.Operation()

}
