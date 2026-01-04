package main

import "fmt"

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("[CPU] Freeze")
}
func (c *CPU) Jump() {
	fmt.Println("[CPU] Jump")
}
func (c *CPU) Execute() {
	fmt.Println("[CPU] Execute")
}

type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("[Memory] Load")
}

type HardDrive struct{}

func (h *HardDrive) Write() {
	fmt.Println("[Hard Drive] Write")
}

type ComputerFascade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func (c *ComputerFascade) Start() {
	c.cpu.Freeze()
	c.cpu.Jump()
	c.memory.Load()
	c.hardDrive.Write()
	c.cpu.Execute()
}

func NewComputerFascade() *ComputerFascade {
	return &ComputerFascade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func main() {
	computer := NewComputerFascade()
	computer.Start()
}
