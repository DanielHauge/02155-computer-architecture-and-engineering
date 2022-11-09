package main

func lw(rd int, rs1 int, imm12 int) {
	reg[rd] = readAsInt(mem[rs1+imm12 : rs1+imm12+4])
}

func lh(rd int, rs1 int, imm12 int) {
	reg[rd] = readAsInt(mem[rs1+imm12 : rs1+imm12+2])
}

func lb(rd int, rs1 int, imm12 int) {
	reg[rd] = readAsInt([]byte{mem[rs1+imm12]})
}

func lwu(rd int, rs1 int, imm12 int) {
	reg[rd] = int(readAsUint(mem[rs1+imm12 : rs1+imm12+4]))
}

func lhu(rd int, rs1 int, imm12 int) {
	reg[rd] = int(readAsUint(mem[rs1+imm12 : rs1+imm12+2]))
}

func lbu(rd int, rs1 int, imm12 int) {
	reg[rd] = int(readAsUint([]byte{mem[rs1+imm12]}))
}

func sw(rs1 int, rs2 int, imm12 int) {
	bytes := castToBytes(reg[rs2])
	mem[rs1+imm12] = bytes[0]
	mem[rs1+imm12+1] = bytes[1]
	mem[rs1+imm12+2] = bytes[2]
	mem[rs1+imm12+3] = bytes[3]
}

func sh(rs1 int, rs2 int, imm12 int) {
	bytes := castToBytes(reg[rs2])
	mem[rs1+imm12] = bytes[0]
	mem[rs1+imm12+1] = bytes[1]
}

func sb(rs1 int, rs2 int, imm12 int) {
	bytes := castToBytes(reg[rs2])
	mem[rs1+imm12] = bytes[0]
}
