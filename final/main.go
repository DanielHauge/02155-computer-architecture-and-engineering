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
	debugMode := os.Getenv("debug") != ""

	binary, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	programLength := int32(len(binary))
	zeros := make([]byte, 100-programLength)
	mem := append(binary, zeros...)
	s.Initialize(mem)
	if debugMode {
		print_memory(len(binary))
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
	print_registers()
}

func print_memory(n int) {
	fmt.Println("\n##### Memory dump")
	for i, b := range s.Mem {
		fmt.Printf("%b ", b)
		if i%4 == 0 {
			fmt.Println()
		}
		if i >= n {
			return
		}
	}
}

func print_registers() {
	fmt.Println()
	for i, r := range s.Reg {
		fmt.Printf(" x%v:\t %v\n", i, r)
	}
}

func dump_registers_to_file(path string) {
	//TODO
}
