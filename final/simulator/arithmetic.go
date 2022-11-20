package simulator

func add(rd int32, rs1 int32, rs2 int32) {
	reg[rd] = reg[rs1] + reg[rs2]
}

func sub(rd int32, rs1 int32, rs2 int32) {
	reg[rd] = reg[rs1] - reg[rs2]
}

func addi(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = reg[rs1] + imm12
}

func slt(rd int32, rs1 int32, rs2 int32) {
	if reg[rs1] < reg[rs2] {
		reg[rd] = 1
	} else {
		reg[rd] = 0
	}
}

func slti(rd int32, rs1 int32, imm12 int32) {
	if reg[rs1] < imm12 {
		reg[rd] = 1
	} else {
		reg[rd] = 0
	}
}

func sltu(rd int32, rs1 int32, rs2 int32) {
	if castToUint(&reg[rs1]) < castToUint(&reg[rs2]) {
		reg[rd] = 1
	} else {
		reg[rd] = 0
	}
}

func sltiu(rd int32, rs1 int32, imm12 int32) {
	if castToUint(&reg[rs1]) < castToUint(&imm12) {
		reg[rd] = 1
	} else {
		reg[rd] = 0
	}
}

func lui(rd int32, imm20 int32) {
	reg[rd] = imm20 << 12
}

func auip(rd int32, imm20 int32) {
	reg[rd] = Pc + imm20<<12
}
