package main

func and(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] & reg[rs2]
}

func or(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] | reg[rs2]
}

func xor(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] ^ reg[rs2]
}

func andi(rd int, rs1 int, imm12 int) {
	reg[rd] = reg[rs1] & imm12
}

func ori(rd int, rs1 int, imm12 int) {
	reg[rd] = reg[rs1] | imm12
}

func xori(rd int, rs1 int, imm12 int) {
	reg[rd] = reg[rs1] ^ imm12
}

// TODO - Shift operations might need a look at.
func sll(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] << reg[rs2]
}

func srl(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] >> reg[rs2]
}

func sra(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] >> reg[rs2]
}

func slli(rd int, rs1 int, imm12 int) {
	reg[rd] = reg[rs1] << imm12
}

func srli(rd int, rs1 int, imm12 int) {
	reg[rd] = reg[rs1] >> imm12
}

func srai(rd int, rs1 int, imm12 int) {
	reg[rd] = reg[rs1] >> imm12
}
