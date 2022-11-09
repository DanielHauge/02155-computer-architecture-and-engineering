package main

func beq(rs1 int, rs2 int, imm12 int) {
	if reg[rs1] == reg[rs2] {
		pc += imm12
	}
}

func bne(rs1 int, rs2 int, imm12 int) {
	if reg[rs1] != reg[rs2] {
		pc += imm12
	}
}

func bge(rs1 int, rs2 int, imm12 int) {
	if reg[rs1] >= reg[rs2] {
		pc += imm12
	}
}

func bgeu(rs1 int, rs2 int, imm12 int) {
	if castToUint(&reg[rs1]) >= castToUint(&reg[rs2]) {
		pc += imm12
	}
}

func blt(rs1 int, rs2 int, imm12 int) {
	if reg[rs1] < reg[rs2] {
		pc += imm12
	}
}

func bltu(rs1 int, rs2 int, imm12 int) {
	if castToUint(&reg[rs1]) < castToUint(&reg[rs2]) {
		pc += imm12
	}
}

func jal(rd int, imm20 int) {
	reg[rd] = pc + 4
	pc = pc + imm20
}

func jalr(rd int, rs1 int, imm12 int) {
	reg[rd] = pc + 4
	pc = reg[rs1] + imm12
}
