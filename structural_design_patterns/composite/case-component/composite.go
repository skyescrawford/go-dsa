package main

type Composite struct {
	components []Component
}

func (c *Composite) Operation() {
	for _, c := range c.components {
		c.Operation()
	}
}

func (c *Composite) add(nc ...Component) {
	for _, v := range nc {
		c.components = append(c.components, v)
	}
}
