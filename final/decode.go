package main

type instruction struct {
	opcode   int
	rd       int
	rs1      int
	rs2      int
	imm12_I  int
	imm12_SB int
	imm20    int
}

func decode(instr []byte) instruction {
	return instruction{}
}

func R(inst instruction, op func(int, int, int)) {
	op(inst.rd, inst.rs1, inst.rs2)
}

func I(inst instruction, op func(int, int, int)) {
	op(inst.rd, inst.rs1, inst.imm12_I)
}

func U(inst instruction, op func(int, int)) {
	op(inst.rd, inst.imm20)
}

func S(inst instruction, op func(int, int, int)) {
	op(inst.rs1, inst.rs2, inst.imm12_SB)
}
