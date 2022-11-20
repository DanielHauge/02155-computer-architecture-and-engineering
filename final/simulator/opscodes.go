package simulator

func wrap3(op func(int32, int32, int32), opFormatFunc func(instruction, func(int32, int32, int32))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

func wrap2(op func(int32, int32), opFormatFunc func(instruction, func(int32, int32))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

func (instr *instruction) Execute() {
	operations[instr.opcode](*instr)
}

var operations = map[int32]func(instruction){
	// Arithmetic
	1: wrap3(add, rFormat),
	2: wrap3(sub, rFormat),
	3: wrap3(addi, iFormat),
	4: wrap3(slt, rFormat),
	5: wrap3(slti, iFormat),
	6: wrap3(sltu, rFormat),
	7: wrap3(sltiu, iFormat),
	8: wrap2(lui, uFormat),
	9: wrap2(auip, uFormat),
	// Logical
	11: wrap3(and, rFormat),
	12: wrap3(or, rFormat),
	13: wrap3(xor, rFormat),
	14: wrap3(andi, iFormat),
	15: wrap3(ori, iFormat),
	16: wrap3(xori, iFormat),
	17: wrap3(sll, rFormat),
	18: wrap3(srl, rFormat),
	19: wrap3(sra, rFormat),
	20: wrap3(slli, iFormat),
	21: wrap3(srli, iFormat),
	22: wrap3(srai, iFormat),
	// Memory
	24: wrap3(lw, iFormat),
	25: wrap3(lh, iFormat),
	26: wrap3(lb, iFormat),
	27: wrap3(lwu, iFormat),
	28: wrap3(lhu, iFormat),
	29: wrap3(lbu, iFormat),
	31: wrap3(sw, sFormat),
	32: wrap3(sh, sFormat),
	33: wrap3(sb, sFormat),
	// Branching
	34: wrap3(beq, sFormat),
	35: wrap3(bne, sFormat),
	36: wrap3(bge, sFormat),
	37: wrap3(bgeu, sFormat),
	38: wrap3(blt, sFormat),
	39: wrap3(bltu, sFormat),
	40: wrap2(jal, uFormat),
	41: wrap3(jalr, iFormat),
}
