package simulator

type instruction struct {
	opcode   int32
	rd       int32
	rs1      int32
	rs2      int32
	imm12_I  int32
	imm12_SB int32
	imm20    int32
}

func Decode(instr []byte) instruction {
	return instruction{
		opcode:   int32(instr[3] & 0x7F),
		rd:       (readAsInt(instr[2:3]) & 0xF80) >> 7,
		rs1:      (readAsInt(instr[1:2]) & 0xF80) >> 7,
		rs2:      (readAsInt(instr[0:1]) & 0x1F0) >> 4,
		imm12_I:  (readAsInt(instr[0:1]) & 0xFFF0) >> 4,
		imm12_SB: ((readAsInt(instr[0:1]) & 0xFE00) >> 2) | ((readAsInt(instr[2:3]) & 0xF80) >> 7),
		imm20:    (readAsInt(instr[0:2]) & 0xFFFFF0) >> 4,
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
