package main

import (
	Cpu "chip8/cpu"
	Rom "chip8/rom"
)

func main() {
	cpu := Cpu.CPU{}
	rom := Rom.ROM{}

	cpu.Init()
	rom.Load("roms/pong.rom", &cpu)
	cpu.LoadRom(rom.Data)
}
