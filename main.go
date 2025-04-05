package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	Cpu "chip8/cpu"
	Graphics "chip8/graphics"
	Rom "chip8/rom"
)

func main() {
	romPath := flag.String("rom", "roms/pong.ch8", "Path to the ROM file")
	delay := flag.Int("delay", 2, "Delay between cycles in milliseconds")
	flag.Parse()

	cpu := Cpu.CPU{}
	rom := Rom.ROM{}
	graphics := Graphics.Graphics{}
	renderer := Graphics.Renderer{}

	cpu.Init()
	if err := rom.Load(*romPath, &cpu); err != nil {
		fmt.Println("Error loading ROM:", err)
		os.Exit(1)
	}

	cpu.LoadRom(rom.Data)
	renderer.Init()

	quit := false

	// main loop
	for !quit {
		quit = renderer.ProcessInput(&cpu.Keypad)

		cpu.Cycle(&graphics)
		renderer.Update(&graphics)
		renderer.PlaySound(cpu.SoundTimer)

		// sleep for the specified delay
		time.Sleep(time.Millisecond * time.Duration(*delay))
	}

	// cleanup resources
	renderer.Cleanup()
}
