package simulator

import (
	"fmt"
	"os"
)

var (
	Pc  int32
	Reg []uint32
	Mem []byte
)

func Initialize(mem []byte) {
	Pc = 0
	Mem = mem
	Reg = make([]uint32, 32)
}

func Print_registers() {
	fmt.Println()
	for i, r := range Reg {
		fmt.Printf(" x%v:\t %X\n", i, r)
	}
}

func Dump_registers_to_file(path string) {
	bs := make([]byte, 4*32)
	for _, i := range Reg {
		b := castUToBytes(i)
		bs = append(bs, b...)
	}
	os.WriteFile(path, bs, 0644)
}

func Print_memory(n int) {
	fmt.Println("\n##### Memory dump")
	for i, b := range Mem {
		fmt.Printf("%b ", b)
		if i%4 == 0 {
			fmt.Println()
		}
		if i >= n {
			return
		}
	}
}

func lw(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = readAsUint(Mem[Reg[rs1]+imm12 : Reg[rs1]+imm12+4])
}

func lh(rd uint32, rs1 uint32, imm12 uint32) {
	hw := readAsHalfWord(Mem[Reg[rs1]+imm12 : Reg[rs1]+imm12+2])
	if (hw >> 15 & 0b1) == 0b1 {
		Reg[rd] = uint32(hw) | 0b11111111111111110000000000000000
	} else {
		Reg[rd] = uint32(hw)
	}
}

func lb(rd uint32, rs1 uint32, imm12 uint32) {
	hw := readAsUint([]byte{Mem[Reg[rs1]+imm12]})
	if (hw >> 7 & 0b1) == 0b1 {
		Reg[rd] = hw | 0b11111111111111111111111100000000
	} else {
		Reg[rd] = hw
	}
}

func lwu(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = readAsUint(Mem[Reg[rs1]+imm12 : Reg[rs1]+imm12+4])
}

func lhu(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = readAsUint(Mem[Reg[rs1]+imm12 : Reg[rs1]+imm12+2])
}

func lbu(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = readAsUint([]byte{Mem[Reg[rs1]+imm12]})
}

func sw(rs1 uint32, rs2 uint32, imm12 uint32) {
	bytes := castUToBytes(Reg[rs2])
	Mem[Reg[rs1]+imm12] = bytes[3]
	Mem[Reg[rs1]+imm12+1] = bytes[2]
	Mem[Reg[rs1]+imm12+2] = bytes[1]
	Mem[Reg[rs1]+imm12+3] = bytes[0]
}

func sh(rs1 uint32, rs2 uint32, imm12 uint32) {
	bytes := castUToBytes(Reg[rs2])
	Mem[Reg[rs1]+imm12] = bytes[1]
	Mem[Reg[rs1]+imm12+1] = bytes[0]
}

func sb(rs1 uint32, rs2 uint32, imm12 uint32) {
	bytes := castUToBytes(Reg[rs2])
	Mem[Reg[rs1]+imm12] = bytes[0]
}
