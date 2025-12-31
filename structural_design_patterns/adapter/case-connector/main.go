package main

import "fmt"

type Computer interface {
	InsertIntoUSBPort()
}

// Adaptee
type LightnightConnector struct{}

func (l *LightnightConnector) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into the Lightning port.")
}

// Adapter
type LightningAdapter struct {
	lightningDevice *LightnightConnector
}

func (l *LightningAdapter) InsertIntoUSBPort() {
	fmt.Println("Adapter converts USB-C signal to Lightning...")
	l.lightningDevice.InsertIntoLightningPort()
}

// Client
type Client struct{}

func (c *Client) ChargeComputer(com Computer) {
	com.InsertIntoUSBPort()
}

func main() {
	client := &Client{}

	lightning := &LightnightConnector{}

	adapter := &LightningAdapter{lightningDevice: lightning}

	client.ChargeComputer(adapter)
}
