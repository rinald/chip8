package cpu

import "chip8/graphics"

func (c *CPU) ExecuteOpcode(g *graphics.Graphics) {
	leftByte := c.opcode >> 8
	rightByte := c.opcode & 0xFF

	switch c.opcode {
	case 0x00E0:
		c.I_00E0(g)
	case 0x00EE:
		c.I_00EE()
	}

	switch leftByte {
	case 0x1:
		c.I_1NNN()
	case 0x2:
		c.I_2NNN()
	case 0x3:
		c.I_3XKK()
	case 0x4:
		c.I_4XKK()
	case 0x5:
		c.I_5XY0()
	case 0x6:
		c.I_6XKK()
	case 0x7:
		c.I_7XKK()
	case 0x8:
		switch rightByte {
		case 0x0:
			c.I_8XY0()
		case 0x1:
			c.I_8XY1()
		case 0x2:
			c.I_8XY2()
		case 0x3:
			c.I_8XY3()
		case 0x4:
			c.I_8XY4()
		case 0x5:
			c.I_8XY5()
		case 0x6:
			c.I_8XY6()
		case 0x7:
			c.I_8XY7()
		case 0xE:
			c.I_8XYE()
		}
	case 0x9:
		c.I_9XY0()
	case 0xA:
		c.I_ANNN()
	case 0xB:
		c.I_BNNN()
	case 0xC:
		c.I_CXKK()
	case 0xD:
		c.I_DXYN(g)
	case 0xE:
		if rightByte == 0x9E {
			c.I_EX9E()
		} else if rightByte == 0xA1 {
			c.I_EXA1()
		}
	case 0xF:
		switch rightByte {
		case 0x07:
			c.I_FX07()
		case 0x0A:
			c.I_FX0A()
		case 0x15:
			c.I_FX15()
		case 0x18:
			c.I_FX18()
		case 0x1E:
			c.I_FX1E()
		case 0x29:
			c.I_FX29()
		case 0x33:
			c.I_FX33()
		case 0x55:
			c.I_FX55()
		case 0x65:
			c.I_FX65()
		}
	}
}
