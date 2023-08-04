package cpu

type Instruction struct {
	Name string
	Exec func(cpu *CPU)
}

type InstructionMap map[byte]Instruction

func ldaSetStatus(cpu *CPU) {
	A := cpu.Registers[A].(*Register8)
	FR := cpu.Registers[FR].(*Register8)
	if A.Value() == 0 {
		FR.SetValue(FR.Value() | 0b01000000)
	}
	if (A.Value() & 0b10000000) > 0 {
		FR.SetValue(FR.Value() | 0b00000001)
	}
}

var INSTRUCTIONS = InstructionMap{
	0x00: Instruction{
		Name: "NOP",
		Exec: func(cpu *CPU) {},
	},
	0xA9: Instruction{
		Name: "LDA_IMM",
		Exec: func(cpu *CPU) {
			imm := cpu.fetch()
			cpu.WriteRegister8(A, imm)
			ldaSetStatus(cpu)
		},
	},
}
