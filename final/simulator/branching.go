package simulator

func beq(rs1 int32, rs2 int32, imm12 int32) {
	if Reg[rs1] == Reg[rs2] {
		Pc += imm12
	}
}

func bne(rs1 int32, rs2 int32, imm12 int32) {
	if Reg[rs1] != Reg[rs2] {
		Pc += imm12
	}
}

func bge(rs1 int32, rs2 int32, imm12 int32) {
	if Reg[rs1] >= Reg[rs2] {
		Pc += imm12
	}
}

func bgeu(rs1 int32, rs2 int32, imm12 int32) {
	if castToUint(&Reg[rs1]) >= castToUint(&Reg[rs2]) {
		Pc += imm12
	}
}

func blt(rs1 int32, rs2 int32, imm12 int32) {
	if Reg[rs1] < Reg[rs2] {
		Pc += imm12
	}
}

func bltu(rs1 int32, rs2 int32, imm12 int32) {
	if castToUint(&Reg[rs1]) < castToUint(&Reg[rs2]) {
		Pc += imm12
	}
}

func jal(rd int32, imm20 int32) {
	Reg[rd] = Pc + 4
	Pc = Pc + imm20
}

func jalr(rd int32, rs1 int32, imm12 int32) {
	Reg[rd] = Pc + 4
	Pc = Reg[rs1] + imm12
}
