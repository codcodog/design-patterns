package facade

import "fmt"

type CPU struct {
}

func (c *CPU) start() string {
	return "cpu start"
}

type Memory struct {
}

func (m *Memory) load() string {
	return "memory load"
}

type HardDrive struct {
}

func (h *HardDrive) read() string {
	return "hard drive read"
}

// facade
type Computer struct {
}

func (c *Computer) Start() string {
	cpu := &CPU{}
	memory := &Memory{}
	hardDrive := &HardDrive{}

	return fmt.Sprintf("%s %s %s",
		cpu.start(),
		memory.load(),
		hardDrive.read())
}
