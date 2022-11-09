package main

func wrap3(op func(int, int, int), opFormatFunc func(instruction, func(int, int, int))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

func wrap2(op func(int, int), opFormatFunc func(instruction, func(int, int))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

var operations = map[int]func(instruction){
	// Arithmetic
	1: wrap3(add, R),
	2: wrap3(sub, R),
	3: wrap3(addi, I),
	4: wrap3(slt, R),
	5: wrap3(slti, I),
	6: wrap3(sltu, R),
	7: wrap3(sltiu, I),
	8: wrap2(lui, U),
	9: wrap2(auip, U),
	// Logical
	11: wrap3(and, R),
	12: wrap3(or, R),
	13: wrap3(xor, R),
	14: wrap3(andi, I),
	15: wrap3(ori, I),
	16: wrap3(xori, I),
	17: wrap3(sll, R),
	18: wrap3(srl, R),
	19: wrap3(sra, R),
	20: wrap3(slli, I),
	21: wrap3(srli, I),
	22: wrap3(srai, I),
	// Memory
	24: wrap3(lw, I),
	25: wrap3(lh, I),
	26: wrap3(lb, I),
	27: wrap3(lwu, I),
	28: wrap3(lhu, I),
	29: wrap3(lbu, I),
	31: wrap3(sw, S),
	32: wrap3(sh, S),
	33: wrap3(sb, S),
	// Branching
	34: wrap3(beq, S),
	35: wrap3(bne, S),
	36: wrap3(bge, S),
	37: wrap3(bgeu, S),
	38: wrap3(blt, S),
	39: wrap3(bltu, S),
	40: wrap2(jal, U),
	41: wrap3(jalr, I),
}
