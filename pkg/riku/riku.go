package riku

import (
	"math"

	"github.com/salty-max/go-riku/pkg/cpu"
	"github.com/salty-max/go-riku/pkg/memory"
	"github.com/salty-max/go-riku/pkg/util"
)

type Riku struct {
	Cpu                  *cpu.CPU
	MM                   *memory.MemoryMapper
	clockSpeed           int
	vBlankInterruptCycle float64
}

func NewRiku() *Riku {
	riku := &Riku{}
	riku.MM = memory.NewMemoryMapper()
	riku.Cpu = cpu.NewCPU(riku.MM)
	riku.clockSpeed = 4194304
	riku.vBlankInterruptCycle = math.Floor(float64(riku.clockSpeed) / 60)

	return riku
}

func (riku *Riku) Init() {
	riku.initMemory()
}

func (riku *Riku) Boot(program []byte) {
	riku.MM.Reset()
	riku.Cpu.Reset()
	riku.loadROM(program)
	riku.loop()
}

func (riku *Riku) initMemory() {
	riku.MM.Map("ZERO_PAGE", memory.CreateRAM(util.ZP_SIZE), util.ZP_START, util.ZP_SIZE, true)
	riku.MM.Map("STACK", memory.CreateRAM(util.STACK_SIZE), util.STACK_END, util.STACK_SIZE, true)
	riku.MM.Map("IO", memory.CreateRAM(util.IO_SIZE), util.IO_START, util.IO_SIZE, true)
	riku.MM.Map("ROM", memory.CreateROM(util.ROM_SIZE), util.ROM_START, util.ROM_SIZE, true)
	riku.MM.Map("BANK_ROM", memory.CreateROM(util.BANK_ROM_SIZE), util.BANK_ROM_START, util.BANK_ROM_SIZE, true)
	riku.MM.Map("CART_DATA", memory.CreateRAM(util.CART_DATA_SIZE), util.CART_DATA_START, util.CART_DATA_SIZE, true)
	riku.MM.Map("RAM", memory.CreateRAM(util.RAM_SIZE), util.RAM_START, util.RAM_SIZE, true)
	riku.MM.Map("VRAM", memory.CreateRAM(util.VRAM_SIZE), util.VRAM_START, util.VRAM_SIZE, true)
	riku.MM.Map("SPRITE_TABLE", memory.CreateRAM(util.SPRITE_TABLE_SIZE), util.SPRITE_TABLE_START, util.SPRITE_TABLE_SIZE, true)
	riku.MM.Map("SID", memory.CreateRAM(util.SID_SIZE), util.SID_START, util.SID_SIZE, true)
	riku.MM.Map("KERNEL", memory.CreateRAM(util.KERNEL_SIZE), util.KERNEL_START, util.KERNEL_SIZE, true)
}

func (riku *Riku) loadROM(program []byte) {
	riku.MM.Load(program, util.ROM_START)
}

func (riku *Riku) loop() {
	riku.Cpu.Cycle()
}
