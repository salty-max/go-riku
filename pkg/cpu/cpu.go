package cpu

import (
	"fmt"

	"github.com/salty-max/go-riku/pkg/memory"
	"github.com/salty-max/go-riku/pkg/util"
)

type CPU struct {
	Memory           memory.Memory
	Registers        RegisterMap
	stack_frame_size int
}

func NewCPU(memory memory.Memory) *CPU {
	cpu := &CPU{Memory: memory, Registers: CreateRegisters(), stack_frame_size: 0}
	cpu.initRegisters()
	return cpu
}

func (cpu *CPU) ViewMemoryAt(address uint16, n int) {
	// 0x0f01: 0x04 0x05 0xA3 0xFE 0x13 0x0D 0x44 0x0F ...
	nextNBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		nextNBytes[i] = cpu.Memory.Read(address + uint16(i))
	}

	fmt.Printf("memory @ %04X: ", address)
	for _, b := range nextNBytes {
		fmt.Printf("%02X ", b)
	}
	fmt.Println()
}

func (c *CPU) initRegisters() {
	c.Registers[R0].(*Register8).Reset()
	c.Registers[RX].(*Register8).Reset()
	c.Registers[RY].(*Register8).Reset()
	c.Registers[A].(*Register8).Reset()
	c.Registers[FR].(*Register8).Reset()
	c.Registers[IR].(*Register8).Reset()
	c.Registers[PC].(*Register16).SetValue(util.ROM_START)
	c.Registers[SP].(*Register16).SetValue(util.STACK_START - 1)
	c.Registers[FP].(*Register16).SetValue(util.STACK_START - 1)
	c.Registers[IE].(*Register8).Reset()
	c.Registers[IF].(*Register8).Reset()
}

func (c *CPU) ReadRegister8(reg Register) uint8 {
	return c.Registers[reg].(*Register8).Value()
}

func (c *CPU) WriteRegister8(reg Register, value uint8) {
	c.Registers[reg].(*Register8).SetValue(value)
}

func (c *CPU) ReadRegister16(reg Register) uint16 {
	return c.Registers[reg].(*Register16).Value()
}

func (c *CPU) WriteRegister16(reg Register, value uint16) {
	c.Registers[reg].(*Register16).SetValue(value)
}

func (c *CPU) fetch() uint8 {
	pc := c.ReadRegister16(PC)
	c.WriteRegister16(PC, pc+1)
	return c.Memory.Read(pc)
}

func (c *CPU) fetch16() uint16 {
	pc := c.ReadRegister16(PC)
	c.WriteRegister16(PC, pc+2)
	return c.Memory.Read16(pc)
}

func (c *CPU) fetchInstruction() Instruction {
	opcode := c.fetch()
	return INSTRUCTIONS[opcode]
}

func (c *CPU) push(value uint16) {
	sp := c.ReadRegister16(SP)
	c.Memory.Write16(sp, value)
	c.WriteRegister16(SP, sp-2)
	c.stack_frame_size += 2
}

func (c *CPU) pop() uint16 {
	sp := c.ReadRegister16(SP)
	c.WriteRegister16(SP, sp+2)
	c.stack_frame_size -= 2
	return c.Memory.Read16(sp)
}

func (c *CPU) Cycle() {
	instruction := c.fetchInstruction()
	instruction.Exec(c)
}

func (c *CPU) Reset() {
	c.initRegisters()
	c.stack_frame_size = 0
}
