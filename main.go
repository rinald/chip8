package main

import (
	"time"

	Cpu "chip8/cpu"
	Graphics "chip8/graphics"
	Rom "chip8/rom"
)

const DELAY = 2 // milliseconds

func main() {
	cpu := Cpu.CPU{}
	rom := Rom.ROM{}
	graphics := Graphics.Graphics{}
	renderer := Graphics.Renderer{}

	cpu.Init()
	rom.Load("roms/pong.ch8", &cpu)
	cpu.LoadRom(rom.Data)
	renderer.Init()

	// main loop
	for {
		cpu.Cycle(&graphics)
		renderer.Update(&graphics)
		time.Sleep(time.Millisecond * DELAY)
	}
}
