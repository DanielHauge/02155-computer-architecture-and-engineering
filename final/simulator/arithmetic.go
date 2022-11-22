package simulator

func add(rd uint32, rs1 uint32, rs2 uint32) {
	Reg[rd] = Reg[rs1] + Reg[rs2]
}

func sub(rd uint32, rs1 uint32, rs2 uint32) {
	Reg[rd] = Reg[rs1] - Reg[rs2]
}

func addi(rd uint32, rs1 uint32, imm12 uint32) {
	Reg[rd] = Reg[rs1] + imm12
}

func slt(rd uint32, rs1 uint32, rs2 uint32) {
	if castToInt(&Reg[rs1]) < castToInt(&Reg[rs2]) {
		Reg[rd] = 1
	} else {
		Reg[rd] = 0
	}
}

func slti(rd uint32, rs1 uint32, imm12 uint32) {
	if castToInt(&Reg[rs1]) < int32(imm12) {
		Reg[rd] = 1
	} else {
		Reg[rd] = 0
	}
}

func sltu(rd uint32, rs1 uint32, rs2 uint32) {
	if Reg[rs1] < Reg[rs2] {
		Reg[rd] = 1
	} else {
		Reg[rd] = 0
	}
}

func sltiu(rd uint32, rs1 uint32, imm12 uint32) {
	if Reg[rs1] < imm12 {
		Reg[rd] = 1
	} else {
		Reg[rd] = 0
	}
}

func lui(rd uint32, imm20 uint32) {
	Reg[rd] = imm20 << 12
}

func auip(rd uint32, imm20 uint32) {
	newRd := uint32(Pc+int32(imm20)) << 12
	Reg[rd] = newRd
}
