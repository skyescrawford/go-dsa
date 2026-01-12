package main

import "fmt"

type Android struct {
	cableData string
}

func (a *Android) plugToUSBPort() {
	fmt.Println("plugging cable data into USB port")
}

type Pluggable interface {
	plugToACSocket()
}

func Charge(p Pluggable) {
	p.plugToACSocket()
}

type PowerAdapter struct {
	phone *Android
}

func (p *PowerAdapter) plugToACSocket() {
	p.phone.plugToUSBPort()
	fmt.Println("plugging power adapter into AC socket")
	fmt.Println("charging...")
}

func main() {
	phone := &Android{}

	adapter := &PowerAdapter{phone}

	Charge(adapter)
}
