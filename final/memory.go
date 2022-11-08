package main

func ld(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = castToInt(mem[rs1+imm12 : rs1+imm12+4])
}

func sd(rs1 int32, rs2 int32, imm12 int32) {
	bytes := castToBytes(reg[rs2])
	mem[rs1+imm12] = bytes[0]
	mem[rs1+imm12+1] = bytes[1]
	mem[rs1+imm12+2] = bytes[2]
	mem[rs1+imm12+3] = bytes[3]
}
