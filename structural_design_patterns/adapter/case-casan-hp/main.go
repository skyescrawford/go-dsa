package main

import "fmt"

// HP usb type C -> casan HP (usb type A) -> (usb port type A) Adapter (steker)  -> stop kontak (steker port)

type IUSBA interface {
	useIUSBA()
}

type ISteker interface {
	useSteker()
}

type CasanHP struct {
	name string
}

func (c CasanHP) useIUSBA() {
	fmt.Println(c.name, "kepala USB tipe A")
}

type Steker struct{}

func (s Steker) useSteker() {
	fmt.Println("Kepala Steker")
}

type usbaToStekerAdapter struct {
	device IUSBA
}

func newAdapter(d IUSBA) usbaToStekerAdapter {
	return usbaToStekerAdapter{device: d}
}

func (a usbaToStekerAdapter) useSteker() {
	fmt.Println("[Adapter] menghubungkan USB tipe A ke Steker")
	a.device.useIUSBA()
	fmt.Println("[Adapter] tersambung ke steker 🚀🚀🚀")
}

func main() {
	casan := CasanHP{name: "casan HP Realll"}
	adapter := newAdapter(casan)
	adapter.useSteker()
}
