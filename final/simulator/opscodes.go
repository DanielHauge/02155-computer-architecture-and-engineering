package simulator

import (
	"fmt"
)

var ECALL = false

func wrap3(op func(uint32, uint32, uint32), opFormatFunc func(instruction, func(uint32, uint32, uint32))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

func wrap2(op func(uint32, uint32), opFormatFunc func(instruction, func(uint32, uint32))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

func (instr *instruction) Execute() {
	op := operation(instr.opcode, instr.funct3, instr.funct7)
	op(*instr)
}

func print_debug(instruction string) {
	if Debug {
		fmt.Printf("%s ", instruction)
	}
}

// https://github.com/michaeljclark/rv8/blob/master/doc/pdf/riscv-instructions.pdf
func operation(opcode uint32, funct3 uint32, funct7 uint32) func(instruction) {
	switch opcode {
	case 0b0110111:
		print_debug("lui")
		return wrap2(lui, uFormat)
	case 0b0010111:
		print_debug("auip")
		return wrap2(auip, uFormat)
	case 0b1101111:
		print_debug("jal")
		return wrap2(jal, jFormat)
	case 0b1100111:
		print_debug("jalr")
		return wrap3(jalr, iFormat)
	case 0b1100011:
		switch funct3 {
		case 0:
			print_debug("beq")
			return wrap3(beq, bFormat)
		case 1:
			print_debug("bne")
			return wrap3(bne, bFormat)
		case 4:
			print_debug("blt")
			return wrap3(blt, bFormat)
		case 5:
			print_debug("bge")
			return wrap3(bge, bFormat)
		case 6:
			print_debug("bltu")
			return wrap3(bltu, bFormat)
		case 7:
			print_debug("bgeu")
			return wrap3(bgeu, bFormat)
		}
	case 0b0000011:
		switch funct3 {
		case 0:
			print_debug("lb")
			return wrap3(lb, iFormat)
		case 1:
			print_debug("lh")
			return wrap3(lh, iFormat)
		case 2:
			print_debug("lw")
			return wrap3(lw, iFormat)
		case 4:
			print_debug("lbu")
			return wrap3(lbu, iFormat)
		case 5:
			print_debug("lhu")
			return wrap3(lhu, iFormat)
		case 6:
			print_debug("lwu")
			return wrap3(lwu, iFormat)
		}
	case 0b0100011:
		switch funct3 {
		case 0:
			print_debug("sb")
			return wrap3(sb, sFormat)
		case 1:
			print_debug("sh")
			return wrap3(sh, sFormat)
		case 2:
			print_debug("sw")
			return wrap3(sw, sFormat)
		}
	case 0b0010011:
		switch funct3 {
		case 0:
			print_debug("addi")
			return wrap3(addi, iFormat)
		case 1:
			print_debug("slli")
			return wrap3(slli, iFormat)
		case 2:
			print_debug("slti")
			return wrap3(slti, iFormat)
		case 3:
			print_debug("sltiu")
			return wrap3(sltiu, iFormat)
		case 4:
			print_debug("xori")
			return wrap3(xori, iFormat)
		case 0b101:
			switch funct7 {
			case 0:
				print_debug("srli")
				return wrap3(srli, iFormat)
			case 0b0100000:
				print_debug("srai")
				return wrap3(srai, iFormat)
			}
		case 6:
			print_debug("ori")
			return wrap3(ori, iFormat)
		case 7:
			print_debug("andi")
			return wrap3(andi, iFormat)
		}
	case 0b0110011: //
		switch funct3 {
		case 0:
			switch funct7 {
			case 0:
				print_debug("add")
				return wrap3(add, rFormat)
			case 0b0100000:
				print_debug("sub")
				return wrap3(sub, rFormat)
			}
		case 1:
			print_debug("sll")
			return wrap3(sll, rFormat)
		case 2:
			print_debug("slt")
			return wrap3(slt, rFormat)
		case 3:
			print_debug("sltu")
			return wrap3(sltu, rFormat)
		case 4:
			print_debug("xor")
			return wrap3(xor, rFormat)
		case 5:
			switch funct7 {
			case 0:
				print_debug("srl")
				return wrap3(srl, rFormat)
			case 0b0100000:
				print_debug("sra")
				return wrap3(sra, rFormat)
			}
		case 6:
			print_debug("or")
			return wrap3(or, rFormat)
		case 7:
			print_debug("and")
			return wrap3(and, rFormat)
		}
	case 0b1110011: //
		return func(i instruction) {
			print_debug("ECALL")
			ECALL = true
		}
	}

	return func(i instruction) {
		fmt.Printf("Could not find operation? opcode: %07b | funct3: %03b | funct7: %07b\n", i.opcode, i.funct3, i.funct7)
	}
}
