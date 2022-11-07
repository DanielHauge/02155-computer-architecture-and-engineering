package main

type instruction struct {
	opcode   int
	rd       int
	rs1      int
	rs2      int
	func3    int
	imm4_0   int
	imm4_1   int
	imm12    int
	imm10_5  int
	imm11    int
	imm11_0  int
	imm11_5  int
	imm20    int
	imm10_1  int
	imm19_12 int
	imm31_12 int
}

func decode(instr []byte) instruction {
	return instruction{}
}

func R(inst instruction, op func(int, int, int)) {
	op(inst.rd, inst.rs1, inst.rs2)
}

func I(inst instruction, op func(int, int, int)) {
	op(inst.rd, inst.rs1, inst.imm19_12)
}
