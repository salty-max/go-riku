package util

const (
	ZP_START                = 0x0000
	ZP_SIZE                 = 0x100
	STACK_START             = 0x01ff
	STACK_SIZE              = 0x100
	STACK_END               = STACK_START - STACK_SIZE + 1
	IO_START                = 0x0200
	IO_SIZE                 = 0x100
	ROM_START               = 0x0300
	ROM_SIZE                = 0x4000
	ROM_END                 = ROM_START + ROM_SIZE - 1
	BANK_ROM_START          = 0x4300
	BANK_ROM_SIZE           = 0x2000
	CART_DATA_START         = 0x6300
	CART_DATA_SIZE          = 0x400
	RAM_START               = 0x6700
	RAM_SIZE                = 0x6000
	RAM_END                 = RAM_START + RAM_SIZE - 1
	VRAM_START              = 0xc700
	VRAM_SIZE               = 0x400
	SPRITE_TABLE_START      = 0xcb00
	SPRITE_TABLE_SIZE       = 0x200
	SID_START               = 0xcd00
	SID_SIZE                = 0x1000
	FREE_2_START            = 0xdd00
	FREE_2_SIZE             = 0x300
	KERNEL_START            = 0xe000
	KERNEL_SIZE             = 0x2000
	INTERRUPT_VECTOR_OFFSET = 0xfe00
	VBLANK_INTERRUPT        = 0xfe00
	IF_OFFSET               = 0xfffa
	IE_OFFSET               = 0xfffb
)
