package cpu

import (
	font "chip8/data"
	"chip8/graphics"
)

type CPU struct {
	registers      [16]byte   // 16 8-bit registers
	memory         [4096]byte // 4KB of memory
	index          uint16     // 16-bit index register
	programCounter uint16     // 16-bit program counter
	stack          [16]uint16 // 16-level stack
	stackPointer   byte       // 8-bit stack pointer
	delayTimer     byte       // delay timer
	opcode         uint16     // current opcode
	SoundTimer     byte       // sound timer
	Keypad         [16]bool   // 16-key keypad
}

// initialize the CPU
func (c *CPU) Init() {
	c.registers = [16]byte{}
	c.memory = [4096]byte{}
	c.index = 0
	c.programCounter = 0
	c.stack = [16]uint16{}
	c.stackPointer = 0
	c.delayTimer = 0
	c.opcode = 0
	c.SoundTimer = 0
	c.Keypad = [16]bool{}
}

// load rom into memory
func (c *CPU) LoadRom(data []byte) {
	// load font data, starting at address 0x50
	for i, b := range font.Font {
		c.memory[0x50+i] = b
	}

	// load rom data into memory, starting at address 0x200
	for i, b := range data {
		c.memory[0x200+i] = b
	}

	// set program counter to start of rom
	c.programCounter = 0x200
}

// cpu cycle
func (c *CPU) Cycle(g *graphics.Graphics) {
	// load opcode and increment program counter
	c.opcode = uint16(c.memory[c.programCounter])<<8 | uint16(c.memory[c.programCounter+1])
	c.programCounter += 2

	// execute current instruction
	c.ExecuteOpcode(g)

	// decrease timers
	if c.delayTimer > 0 {
		c.delayTimer--
	}
	if c.SoundTimer > 0 {
		c.SoundTimer--
	}
}
