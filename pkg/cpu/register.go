package cpu

type Register int

const (
	R0 Register = iota
	RX
	RY
	A
	FR
	IR
	PC
	SP
	FP
	IE
	IF
)

type Register8 struct {
	value uint8
}

func NewRegister8(v uint8) *Register8 {
	return &Register8{value: v & 0xff}
}

func (r *Register8) Value() uint8 {
	return r.value
}

func (r *Register8) SetValue(v uint8) {
	r.value = v & 0xff
}

func (r *Register8) Reset() {
	r.value = 0
}

type Register16 struct {
	value uint16
}

func NewRegister16(v uint16) *Register16 {
	return &Register16{value: v & 0xffff}
}

func (r *Register16) Value() uint16 {
	return r.value
}

func (r *Register16) SetValue(v uint16) {
	r.value = v & 0xffff
}

func (r *Register16) Reset() {
	r.value = 0
}

type RegisterMap map[Register]interface{}

func CreateRegisters() RegisterMap {
	return RegisterMap{
		R0: NewRegister8(0),
		RX: NewRegister8(0),
		RY: NewRegister8(0),
		A:  NewRegister8(0),
		FR: NewRegister8(0),
		IR: NewRegister8(0),
		PC: NewRegister16(0),
		SP: NewRegister16(0),
		FP: NewRegister16(0),
		IE: NewRegister8(0),
		IF: NewRegister8(0),
	}
}
