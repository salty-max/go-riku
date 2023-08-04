package memory

import "fmt"

type Memory interface {
	Read(offset uint16) byte
	Read16(offset uint16) uint16
	Write(offset uint16, value byte)
	Write16(offset uint16, value uint16)
	Load(buffer []byte, offset uint16)
	Slice(start uint16, end uint16) []byte
	Reset()
}

type RAM struct {
	Buffer []byte
}

func (r *RAM) Read(offset uint16) byte {
	return r.Buffer[offset]
}

func (r *RAM) Read16(offset uint16) uint16 {
	return (uint16(r.Buffer[offset+1]) << 8) | uint16(r.Buffer[offset])
}

func (r *RAM) Write(offset uint16, value byte) {
	r.Buffer[offset] = value
}

func (r *RAM) Write16(offset uint16, value uint16) {
	r.Buffer[offset] = byte(value & 0xFF)
	r.Buffer[offset+1] = byte(value >> 8)
}

func (r *RAM) Load(buffer []byte, offset uint16) {
	copy(r.Buffer, buffer[offset:])
}

func (r *RAM) Slice(start uint16, end uint16) []byte {
	return r.Buffer[start:end]
}

func (r *RAM) Reset() {
	for i := 0; i < len(r.Buffer); i++ {
		r.Buffer[i] = 0
	}
}

func CreateRAM(size uint16) Memory {
	return &RAM{Buffer: make([]byte, size)}
}

type ROM struct {
	Buffer []byte
}

func (r *ROM) Read(offset uint16) byte {
	return r.Buffer[offset]
}

func (r *ROM) Read16(offset uint16) uint16 {
	return (uint16(r.Buffer[offset+1]) << 8) | uint16(r.Buffer[offset])
}

func (r *ROM) Write(offset uint16, value byte) {
	fmt.Println("Cannot write to ROM")
}
func (r *ROM) Write16(offset uint16, value uint16) {
	fmt.Println("Cannot write to ROM")
}

func (r *ROM) Load(buffer []byte, offset uint16) {
	copy(r.Buffer, buffer[offset:])
}

func (r *ROM) Slice(start uint16, end uint16) []byte {
	return r.Buffer[start:end]
}

func (r *ROM) Reset() {
	for i := 0; i < len(r.Buffer); i++ {
		r.Buffer[i] = 0
	}
}

func CreateROM(size uint16) Memory {
	return &ROM{Buffer: make([]byte, size)}
}
