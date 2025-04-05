package main

import (
	Cpu "chip8/cpu"
	Graphics "chip8/graphics"
	Rom "chip8/rom"
)

func main() {
	cpu := Cpu.CPU{}
	rom := Rom.ROM{}
	graphics := Graphics.Graphics{}
	renderer := Graphics.Renderer{}

	cpu.Init()
	rom.Load("roms/pong.ch8", &cpu)
	cpu.LoadRom(rom.Data)
	renderer.Init()

	quit := false

	// main loop
	for !quit {
		quit = renderer.ProcessInput(&cpu.Keypad)

		cpu.Cycle(&graphics)
		renderer.Update(&graphics)
		renderer.PlaySound(cpu.SoundTimer)
	}

	// cleanup resources
	renderer.Cleanup()
}
