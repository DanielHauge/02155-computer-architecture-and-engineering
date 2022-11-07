package main

func ld(rd int, rs1 int, imm12 int) {
	reg[rd] = castToInt(mem[rs1+imm12 : rs1+imm12+4])
}

func sd(rs1 int, rs2 int, imm12 int) {
	bytes := castToBytes(reg[rs2])
	mem[rs1+imm12] = bytes[0]
	mem[rs1+imm12+1] = bytes[1]
	mem[rs1+imm12+2] = bytes[2]
	mem[rs1+imm12+3] = bytes[3]
}
