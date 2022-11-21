package simulator

func add(rd int32, rs1 int32, rs2 int32) {
	Reg[rd] = Reg[rs1] + Reg[rs2]
}

func sub(rd int32, rs1 int32, rs2 int32) {
	Reg[rd] = Reg[rs1] - Reg[rs2]
}

func addi(rd int32, rs1 int32, imm12 int32) {
	Reg[rd] = Reg[rs1] + imm12
}

func slt(rd int32, rs1 int32, rs2 int32) {
	if Reg[rs1] < Reg[rs2] {
		Reg[rd] = 1
	} else {
		Reg[rd] = 0
	}
}

func slti(rd int32, rs1 int32, imm12 int32) {
	if Reg[rs1] < imm12 {
		Reg[rd] = 1
	} else {
		Reg[rd] = 0
	}
}

func sltu(rd int32, rs1 int32, rs2 int32) {
	if castToUint(&Reg[rs1]) < castToUint(&Reg[rs2]) {
		Reg[rd] = 1
	} else {
		Reg[rd] = 0
	}
}

func sltiu(rd int32, rs1 int32, imm12 int32) {
	if castToUint(&Reg[rs1]) < castToUint(&imm12) {
		Reg[rd] = 1
	} else {
		Reg[rd] = 0
	}
}

func lui(rd int32, imm20 int32) {
	Reg[rd] = imm20 << 12
}

func auip(rd int32, imm20 int32) {
	Reg[rd] = Pc + imm20<<12
}
