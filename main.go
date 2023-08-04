package main

import (
	"github.com/salty-max/go-riku/pkg/riku"
)

func main() {
	riku := riku.NewRiku()
	riku.Init()

	riku.Boot([]byte{})
}
