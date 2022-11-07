package main

import (
	"fmt"
)

var (
	pc  int
	reg [32]int
	mem [1000000]byte
)

// This is the main entry of the simulator
func main() {
	fmt.Println("Hello, World!")

	for {
		execute(decode(mem[pc : pc+4]))
		pc += 4
	}
}

func execute(inst instruction) {
	switch inst.opcode {
	case 25:
		R(inst, add)
	case 21:
		R(inst, sub)
	case 52:
		I(inst, sd)
	case 51:
		I(inst, ld)
	}
}
