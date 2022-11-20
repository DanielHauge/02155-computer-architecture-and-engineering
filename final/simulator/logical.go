package simulator

func and(rd int32, rs1 int32, rs2 int32) {
	reg[rd] = reg[rs1] & reg[rs2]
}

func or(rd int32, rs1 int32, rs2 int32) {
	reg[rd] = reg[rs1] | reg[rs2]
}

func xor(rd int32, rs1 int32, rs2 int32) {
	reg[rd] = reg[rs1] ^ reg[rs2]
}

func andi(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = reg[rs1] & imm12
}

func ori(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = reg[rs1] | imm12
}

func xori(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = reg[rs1] ^ imm12
}

// TODO - Shift operations might need a look at.
func sll(rd int32, rs1 int32, rs2 int32) {
	reg[rd] = reg[rs1] << reg[rs2]
}

func srl(rd int32, rs1 int32, rs2 int32) {
	reg[rd] = reg[rs1] >> reg[rs2]
}

func sra(rd int32, rs1 int32, rs2 int32) {
	reg[rd] = reg[rs1] >> reg[rs2]
}

func slli(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = reg[rs1] << imm12
}

func srli(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = reg[rs1] >> imm12
}

func srai(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = reg[rs1] >> imm12
}
