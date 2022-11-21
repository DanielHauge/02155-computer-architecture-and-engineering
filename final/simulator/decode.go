package simulator

type instruction struct {
	opcode   int32
	funct3   int32
	funct7   int32
	rd       int32
	rs1      int32
	rs2      int32
	imm12_I  int32
	imm12_SB int32
	imm20    int32
}

// Opcodes from: https://github.com/michaeljclark/rv8/blob/master/doc/pdf/riscv-instructions.pdf
func Decode(instr []byte) instruction {
	// Reverse because of endian order. (I've mistakenly worked with big endian, where instructions come in little endian)
	for i, j := 0, len(instr)-1; i < j; i, j = i+1, j-1 {
		instr[i], instr[j] = instr[j], instr[i]
	}
	return instruction{
		opcode:   (readAsInt(instr) & 0x7F),
		funct3:   (readAsInt(instr) >> 12) & 7,
		funct7:   (readAsInt(instr) >> 25) & 0x7F,
		rd:       (readAsInt(instr) >> 7) & 0b11111,
		rs1:      (readAsInt(instr) >> 15) & 0b11111,
		rs2:      (readAsInt(instr) >> 20) & 0b11111,
		imm12_I:  ((readAsInt(instr) & -1048576) >> 20) & 0b111111111111,
		imm12_SB: ((readAsInt(instr) >> 20) & 0xFE0) | ((readAsInt(instr) >> 7) & 0x1F),
		imm20:    ((readAsInt(instr) & -4096) >> 12) & 0xFFFFF,
	}
}

func rFormat(inst instruction, op func(int32, int32, int32)) {
	op(inst.rd, inst.rs1, inst.rs2)
}

func iFormat(inst instruction, op func(int32, int32, int32)) {
	op(inst.rd, inst.rs1, inst.imm12_I)
}

func uFormat(inst instruction, op func(int32, int32)) {
	op(inst.rd, inst.imm20)
}

func sFormat(inst instruction, op func(int32, int32, int32)) {
	op(inst.rs1, inst.rs2, inst.imm12_SB)
}
