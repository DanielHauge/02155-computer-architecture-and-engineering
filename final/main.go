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
		inst := decode(mem[pc : pc+4]) // Decode instruction from binary
		operations[inst.opcode](inst)  // Execute operation from instruction
		pc += 4                        // Increment program counter
	}
}
