package simulator

func beq(rs1 uint32, rs2 uint32, imm12 uint32) {
	if castToInt(&Reg[rs1]) == castToInt(&Reg[rs2]) {
		Pc += castToInt(&imm12) - 4 // -4 to only mark those operations that should not increment PC by 4 by default.
	}
}

func bne(rs1 uint32, rs2 uint32, imm12 uint32) {
	if castToInt(&Reg[rs1]) != castToInt(&Reg[rs2]) {
		Pc += castToInt(&imm12) - 4 // -4 to only mark those operations that should not increment PC by 4 by default.
	}
}

func bge(rs1 uint32, rs2 uint32, imm12 uint32) {
	if castToInt(&Reg[rs1]) >= castToInt(&Reg[rs2]) {
		Pc += castToInt(&imm12) - 4 // -4 to only mark those operations that should not increment PC by 4 by default.
	}
}

func bgeu(rs1 uint32, rs2 uint32, imm12 uint32) {
	if Reg[rs1] >= Reg[rs2] {
		Pc += int32(imm12) - 4 // -4 to only mark those operations that should not increment PC by 4 by default.
	}
}

func blt(rs1 uint32, rs2 uint32, imm12 uint32) {
	if castToInt(&Reg[rs1]) < castToInt(&Reg[rs2]) {
		Pc += castToInt(&imm12) - 4 // -4 to only mark those operations that should not increment PC by 4 by default.
	}
}

func bltu(rs1 uint32, rs2 uint32, imm12 uint32) {
	if Reg[rs1] < Reg[rs2] {
		Pc += int32(imm12) - 4 // -4 to only mark those operations that should not increment PC by 4 by default.
	}
}

func jal(rd uint32, imm20 uint32) {
	Reg[rd] = uint32(Pc + 4)
	Pc = Pc + int32(imm20)
}

func jalr(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = uint32(Pc + 4)
	Pc = castToInt(&Reg[rs1]) + castToInt(&imm12)
}
