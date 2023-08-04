package main

import (
	"fmt"

	"github.com/salty-max/go-riku/pkg/riku"
)

func main() {
	riku := riku.NewRiku()
	riku.Init()
	riku.Boot([]byte{
		0xA9, 0x84, // LDA_IMM 0x42
	})

	fmt.Printf("A: %d", riku.Cpu.ReadRegister8(3))
}
