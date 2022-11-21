package main

import (
	s "final/simulator"
	"fmt"
	"os"
)

// This is the main entry of the simulator
func main() {
	if len(os.Args) < 2 {
		fmt.Print("Expected a filepath as argument, but received none.\n")
		os.Exit(1)
	}
	fmt.Printf("Simulating program from binaries in: %s\n", os.Args[1])
	debugMode := os.Getenv("debug") == ""

	binary, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	programLength := int32(len(binary))
	zeros := make([]byte, 100000-programLength)
	mem := append(binary, zeros...)
	s.Initialize(mem)
	if debugMode {
		s.Print_memory(len(binary))
		fmt.Println()
	}

	for s.Pc < programLength && !s.ECALL {
		instrBs := s.Mem[s.Pc : s.Pc+4] // Fetch instructions from memory
		if debugMode {
			fmt.Printf("Executing: %v \n", instrBs)
		}
		inst := s.Decode(instrBs) // Decode instruction from binary
		inst.Execute()            // Execute operation from instruction
		s.Pc += 4                 // Increment program counter
	}
	s.Print_registers()
	if len(os.Args) > 2 {
		s.Dump_registers_to_file(os.Args[2])
	}
}
