package main

import "fmt"

// misalnya user punya hp Iphone 17 pro max, yang kabel chargerya USB C to USB C
// karena ga ada colokan listrik dengan port USB C maka perlu adapter
// adapter ini memiliki soket USB C dan colokan standar yg kompatibel dengan stopkontak pada umumnya
// jadi spesifikasi adapter adalah, menerima USB C plug dan colokan standar
// mari kita bikin

type Iphone struct{}

type USBCtoElectricalPlugAdapter struct {
	phone *Iphone
}

func (u *USBCtoElectricalPlugAdapter) plug() {
	fmt.Println("Plugging USB C to Power Outlet..")
}

type Pluggable interface {
	plug()
}

func charge(c Pluggable) {
	c.plug()
	fmt.Println("Charging... 🚀")
}

func main() {

	iphone17promax := &Iphone{}

	adapter := &USBCtoElectricalPlugAdapter{phone: iphone17promax}

	charge(adapter)

}
