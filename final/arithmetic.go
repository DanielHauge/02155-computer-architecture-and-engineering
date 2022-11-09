package main

func add(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] + reg[rs2]
}

func sub(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] - reg[rs2]
}

func addi(rd int, rs1 int, imm12 int) {
	reg[rd] = reg[rs1] + imm12
}

func slt(rd int, rs1 int, rs2 int) {
	if reg[rs1] < reg[rs2] {
		reg[rd] = 1
	} else {
		reg[rd] = 0
	}
}

func slti(rd int, rs1 int, imm12 int) {
	if reg[rs1] < imm12 {
		reg[rd] = 1
	} else {
		reg[rd] = 0
	}
}

func sltu(rd int, rs1 int, rs2 int) {
	if castToUint(&reg[rs1]) < castToUint(&reg[rs2]) {
		reg[rd] = 1
	} else {
		reg[rd] = 0
	}
}

func sltiu(rd int, rs1 int, imm12 int) {
	if castToUint(&reg[rs1]) < castToUint(&imm12) {
		reg[rd] = 1
	} else {
		reg[rd] = 0
	}
}

func lui(rd int, imm20 int) {
	reg[rd] = imm20 << 12
}

func auip(rd int, imm20 int) {
	reg[rd] = pc + imm20<<12
}
