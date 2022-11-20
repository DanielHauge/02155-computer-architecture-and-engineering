package main

import (
	s "final/simulator"
	"fmt"
	"os"
)

// This is the main entry of the simulator
func main() {
	fmt.Println("Hello, World!")

	binary, err := os.ReadFile("gg")
	if err != nil {
		panic(err)
	}
	zeros := make([]byte, 10000-len(binary))
	mem := append(binary, zeros...)
	s.Initialize(mem)

	for {
		inst := s.Decode(s.Mem[s.Pc : s.Pc+4]) // Decode instruction from binary
		inst.Execute()                         // Execute operation from instruction
		s.Pc += 4                              // Increment program counter
	}
}
