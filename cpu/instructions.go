package cpu

import (
	"chip8/graphics"
	"math/rand"
)

// clear screen (CLS)
func (c *CPU) I_00E0(g *graphics.Graphics) {
	g.Clear()
}

// return from subroutine (RET)
func (c *CPU) I_00EE() {
	c.stackPointer--
	c.programCounter = c.stack[c.stackPointer]
}

// jump to address NNN (JP addr)
func (c *CPU) I_1NNN() {
	address := c.opcode & 0x0FFF
	c.programCounter = address
}

// call subroutine at NNN (CALL addr)
func (c *CPU) I_2NNN() {
	address := c.opcode & 0x0FFF
	c.stack[c.stackPointer] = c.programCounter
	c.stackPointer++
	c.programCounter = address
}

// skip next instruction if Vx == kk (SE Vx, byte)
func (c *CPU) I_3XKK() {
	Vx := (c.opcode & 0x0F00) >> 8
	kk := c.opcode & 0x00FF
	if c.registers[Vx] == byte(kk) {
		c.programCounter += 2
	}
}

// skip next instruction if Vx != kk (SNE Vx, byte)
func (c *CPU) I_4XKK() {
	Vx := (c.opcode & 0x0F00) >> 8
	kk := c.opcode & 0x00FF
	if c.registers[Vx] != byte(kk) {
		c.programCounter += 2
	}
}

// skip next instruction if Vx == Vy (SE Vx, Vy)
func (c *CPU) I_5XY0() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4
	if c.registers[Vx] == c.registers[Vy] {
		c.programCounter += 2
	}
}

// set Vx = kk (LD Vx, byte)
func (c *CPU) I_6XKK() {
	Vx := (c.opcode & 0x0F00) >> 8
	kk := c.opcode & 0x00FF
	c.registers[Vx] = byte(kk)
}

// set Vx = Vx + kk (ADD Vx, byte)
func (c *CPU) I_7XKK() {
	Vx := (c.opcode & 0x0F00) >> 8
	kk := c.opcode & 0x00FF
	c.registers[Vx] += byte(kk)
}

// set Vx = Vy (LD Vx, Vy)
func (c *CPU) I_8XY0() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4
	c.registers[Vx] = c.registers[Vy]
}

// set Vx = Vx OR Vy (OR Vx, Vy)
func (c *CPU) I_8XY1() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4
	c.registers[Vx] |= c.registers[Vy]
}

// set Vx = Vx AND Vy (AND Vx, Vy)
func (c *CPU) I_8XY2() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4
	c.registers[Vx] &= c.registers[Vy]
}

// set Vx = Vx XOR Vy (XOR Vx, Vy)
func (c *CPU) I_8XY3() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4
	c.registers[Vx] ^= c.registers[Vy]
}

// set Vx = Vx + Vy, set VF = carry (ADD Vx, Vy)
func (c *CPU) I_8XY4() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4
	sum := c.registers[Vx] + c.registers[Vy]

	// set carry flag
	if sum > 0xFF {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}

	c.registers[Vx] = sum & 0xFF
}

// set Vx = Vx - Vy, set VF = NOT borrow (SUB Vx, Vy)
func (c *CPU) I_8XY5() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4

	if c.registers[Vx] > c.registers[Vy] {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}

	c.registers[Vx] -= c.registers[Vy]
}

// set Vx = Vx SHR 1 (SHR Vx)
func (c *CPU) I_8XY6() {
	Vx := (c.opcode & 0x0F00) >> 8
	c.registers[0xF] = c.registers[Vx] & 0x1
	c.registers[Vx] >>= 1
}

// set Vx = Vy - Vx, set VF = NOT borrow (SUBN Vx, Vy)
func (c *CPU) I_8XY7() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4

	if c.registers[Vy] > c.registers[Vx] {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}

	c.registers[Vx] = c.registers[Vy] - c.registers[Vx]
}

// set Vx = Vx SHL 1 (SHL Vx)
func (c *CPU) I_8XYE() {
	Vx := (c.opcode & 0x0F00) >> 8
	c.registers[0xF] = (c.registers[Vx] & 0x80) >> 7
	c.registers[Vx] <<= 1
}

// skip next instruction if Vx != Vy (SNE Vx, Vy)
func (c *CPU) I_9XY0() {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4

	if c.registers[Vx] != c.registers[Vy] {
		c.programCounter += 2
	}
}

// set I = NNN (LD I, addr)
func (c *CPU) I_ANNN() {
	c.index = (c.opcode & 0x0FFF)
}

// jump to location NNN + V0 (JP V0, addr)
func (c *CPU) I_BNNN() {
	c.programCounter = (c.opcode & 0x0FFF) + uint16(c.registers[0])
}

// set Vx = random byte AND kk
func (c *CPU) I_CXKK() {
	Vx := (c.opcode & 0x0F00) >> 8
	kk := byte(c.opcode & 0x00FF)
	c.registers[Vx] = byte(rand.Intn(0xFF)) & kk
}

// display n-byte sprite at (Vx, Vy), set VF = collision (DRW Vx, Vy, nibble)
func (c *CPU) I_DXYN(g *graphics.Graphics) {
	Vx := (c.opcode & 0x0F00) >> 8
	Vy := (c.opcode & 0x00F0) >> 4
	n := (c.opcode & 0x000F)

	x := c.registers[Vx] % graphics.WIDTH
	y := c.registers[Vy] % graphics.HEIGHT

	c.registers[0xF] = 0

	for i := uint16(0); i < n; i++ {
		sprite := c.memory[c.index+i]
		for j := byte(0); j < 8; j++ {
			spritePixel := sprite & (0x80 >> j)
			screenPixel := &g.Screen[(y+byte(i))*graphics.WIDTH+x+j]

			if spritePixel != 0 {
				if *screenPixel == 0xFFFFFFFF {
					c.registers[0xF] = 1
				}
				*screenPixel ^= 0xFFFFFFFF
			}
		}
	}
}

// skip next instruction if key with the value of Vx is pressed (SKP Vx)
func (c *CPU) I_EX9E() {
	Vx := (c.opcode & 0x0F00) >> 8
	if c.keypad[c.registers[Vx]] {
		c.programCounter += 2
	}
}

// skip next instruction if key with the value of Vx is not pressed (SKNP Vx)
func (c *CPU) I_EXA1() {
	Vx := (c.opcode & 0x0F00) >> 8
	if !c.keypad[c.registers[Vx]] {
		c.programCounter += 2
	}
}

// set Vx = delay timer (LD Vx, DT)
func (c *CPU) I_FX07() {
	Vx := (c.opcode & 0x0F00) >> 8
	c.registers[Vx] = c.delayTimer
}

// wait for a key press, store the value of the key in Vx (LD Vx, K)
func (c *CPU) I_FX0A() {
	Vx := (c.opcode & 0x0F00) >> 8
	for i := byte(0); i < 16; i++ {
		if c.keypad[i] {
			c.registers[Vx] = i
			return
		}
	}
	c.programCounter -= 2
}

// set delay timer = Vx (LD DT, Vx)
func (c *CPU) I_FX15() {
	Vx := (c.opcode & 0x0F00) >> 8
	c.delayTimer = c.registers[Vx]
}

// set sound timer = Vx (LD ST, Vx)
func (c *CPU) I_FX18() {
	Vx := (c.opcode & 0x0F00) >> 8
	c.soundTimer = c.registers[Vx]
}

// set I = I + Vx (ADD I, Vx)
func (c *CPU) I_FX1E() {
	Vx := (c.opcode & 0x0F00) >> 8
	c.index += uint16(c.registers[Vx])
}

// set I = location of sprite for digit Vx (LD F, Vx)
func (c *CPU) I_FX29() {
	Vx := (c.opcode & 0x0F00) >> 8
	c.index = 0x50 + uint16(c.registers[Vx])*5
}

// store BCD of Vx in memory locations I, I+1, and I+2 (LD B, Vx)
func (c *CPU) I_FX33() {
	Vx := (c.opcode & 0x0F00) >> 8
	c.memory[c.index] = c.registers[Vx] / 100
	c.memory[c.index+1] = (c.registers[Vx] / 10) % 10
	c.memory[c.index+2] = c.registers[Vx] % 10
}

// store registers V0 through Vx in memory starting at location I (LD [I], Vx)
func (c *CPU) I_FX55() {
	Vx := (c.opcode & 0x0F00) >> 8
	for i := uint16(0); i <= Vx; i++ {
		c.memory[c.index+i] = c.registers[i]
	}
}

// read registers V0 through Vx from memory starting at location I (LD Vx, [I])
func (c *CPU) I_FX65() {
	Vx := (c.opcode & 0x0F00) >> 8
	for i := uint16(0); i <= Vx; i++ {
		c.registers[i] = c.memory[c.index+i]
	}
}
