package main

import "fmt"

type Subject interface {
	Request()
}

type RealSubject struct{}

func (r RealSubject) Request() {
	fmt.Println("handling request by RealSuject")
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	fmt.Println("logging before request")
	p.realSubject.Request()
	fmt.Println("logging after request")
}

func main() {
	proxy := &Proxy{}

	proxy.Request()
}
