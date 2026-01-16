package main

import (
	"log"
	"github.com/ice-rider/go_snake/graphics"
)

func main() {
	p := graphics.NewProgram(15, 30)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}