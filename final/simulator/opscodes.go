package simulator

import "fmt"

var ECALL = false

func wrap3(op func(int32, int32, int32), opFormatFunc func(instruction, func(int32, int32, int32))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

func wrap2(op func(int32, int32), opFormatFunc func(instruction, func(int32, int32))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

func (instr *instruction) Execute() {
	op := operation(instr.opcode, instr.funct3, instr.funct7)
	op(*instr)
}

// https://github.com/michaeljclark/rv8/blob/master/doc/pdf/riscv-instructions.pdf
func operation(opcode int32, funct3 int32, funct7 int32) func(instruction) {
	switch opcode {
	case 0b0110111:
		return wrap2(lui, uFormat)
	case 0b0010111:
		return wrap2(auip, uFormat)
	case 0b1101111:
		return wrap2(jal, uFormat)
	case 0b1100111:
		return wrap3(jalr, iFormat)
	case 0b1100011:
		switch funct3 {
		case 0:
			return wrap3(beq, sFormat)
		case 1:
			return wrap3(bne, sFormat)
		case 4:
			return wrap3(blt, sFormat)
		case 5:
			return wrap3(bge, sFormat)
		case 6:
			return wrap3(bltu, sFormat)
		case 7:
			return wrap3(bgeu, sFormat)
		}
	case 0b0000011:
		switch funct3 {
		case 0:
			return wrap3(lb, iFormat)
		case 1:
			return wrap3(lh, iFormat)
		case 2:
			return wrap3(lw, iFormat)
		case 4:
			return wrap3(lbu, iFormat)
		case 5:
			return wrap3(lhu, iFormat)
		case 6:
			return wrap3(lwu, iFormat)
		}
	case 0b0100011:
		switch funct3 {
		case 0:
			return wrap3(sb, sFormat)
		case 1:
			return wrap3(sh, sFormat)
		case 2:
			return wrap3(sw, sFormat)
		}
	case 0b0010011:
		switch funct3 {
		case 0:
			return wrap3(addi, iFormat)
		case 1:
			return wrap3(slli, iFormat)
		case 2:
			return wrap3(slti, iFormat)
		case 3:
			return wrap3(sltiu, iFormat)
		case 4:
			return wrap3(xori, iFormat)
		case 5:
			switch funct7 {
			case 0:
				return wrap3(srli, iFormat)
			case 0x8:
				return wrap3(srai, iFormat)
			}
		case 6:
			return wrap3(ori, iFormat)
		case 7:
			return wrap3(andi, iFormat)
		}
	case 0b0110011: //
		switch funct3 {
		case 0:
			switch funct7 {
			case 0:
				return wrap3(add, rFormat)
			case 0x8:
				return wrap3(sub, rFormat)
			}
		case 1:
			return wrap3(sll, rFormat)
		case 2:
			return wrap3(slt, rFormat)
		case 3:
			return wrap3(sltu, rFormat)
		case 4:
			return wrap3(xor, rFormat)
		case 5:
			switch funct7 {
			case 0:
				return wrap3(srl, rFormat)
			case 0x8:
				return wrap3(sra, rFormat)
			}
		case 6:
			return wrap3(or, rFormat)
		case 7:
			return wrap3(and, rFormat)
		}
	case 0b1110011: //
		return func(i instruction) { ECALL = true }
	}

	return func(i instruction) {
		fmt.Printf("Could not find operation? opcode: %v\n", i.opcode)
	}
}
