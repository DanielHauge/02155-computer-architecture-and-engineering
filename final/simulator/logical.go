package simulator

import (
	"strconv"
	"strings"
)

func and(rd uint32, rs1 uint32, rs2 uint32) {
	Reg[rd] = Reg[rs1] & Reg[rs2]
}

func or(rd uint32, rs1 uint32, rs2 uint32) {
	Reg[rd] = Reg[rs1] | Reg[rs2]
}

func xor(rd uint32, rs1 uint32, rs2 uint32) {
	Reg[rd] = Reg[rs1] ^ Reg[rs2]
}

func andi(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = Reg[rs1] & imm12
}

func ori(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = Reg[rs1] | imm12
}

func xori(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = Reg[rs1] ^ imm12
}

func sll(rd uint32, rs1 uint32, rs2 uint32) {
	Reg[rd] = Reg[rs1] << Reg[rs2]
}

func srl(rd uint32, rs1 uint32, rs2 uint32) {
	Reg[rd] = Reg[rs1] >> Reg[rs2]
}

func sra(rd uint32, rs1 uint32, rs2 uint32) {
	amount := Reg[rs2]
	var s strings.Builder
	for i := 0; i < int(amount); i++ {
		s.WriteString("1")
	}
	for i := int(amount); i < 32; i++ {
		s.WriteString("0")
	}
	u, _ := strconv.ParseUint(s.String(), 2, 32)

	Reg[rd] = (Reg[rs1] >> Reg[rs2]) | uint32(u)
}

func slli(rd uint32, rs1 uint32, imm12 uint32) {
	imm12 = imm12 & 0b11111
	Reg[rd] = Reg[rs1] << imm12
}

func srli(rd uint32, rs1 uint32, imm12 uint32) {
	imm12 = imm12 & 0b11111
	Reg[rd] = Reg[rs1] >> imm12
}

func srai(rd uint32, rs1 uint32, imm12 uint32) {
	imm12 = imm12 & 0b11111
	var s strings.Builder
	for i := 0; i < int(imm12); i++ {
		s.WriteString("1")
	}
	for i := int(imm12); i < 32; i++ {
		s.WriteString("0")
	}
	u, _ := strconv.ParseUint(s.String(), 2, 32)
	Reg[rd] = (Reg[rs1] >> imm12) | uint32(u)
}
