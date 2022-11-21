package simulator

var (
	Pc  int32
	reg []int32
	Mem []byte
)

func Initialize(mem []byte) {
	Pc = 0
	Mem = mem
	reg = make([]int32, 32)
	reg[2] = int32(len(mem))
}

func lw(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = readAsInt(Mem[reg[rs1]+imm12 : reg[rs1]+imm12+4])
}

func lh(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = readAsInt(Mem[reg[rs1]+imm12 : reg[rs1]+imm12+2])
}

func lb(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = readAsInt([]byte{Mem[reg[rs1]+imm12]})
}

func lwu(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = int32(readAsUint(Mem[reg[rs1]+imm12 : reg[rs1]+imm12+4]))
}

func lhu(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = int32(readAsUint(Mem[reg[rs1]+imm12 : reg[rs1]+imm12+2]))
}

func lbu(rd int32, rs1 int32, imm12 int32) {
	reg[rd] = int32(readAsUint([]byte{Mem[reg[rs1]+imm12]}))
}

func sw(rs1 int32, rs2 int32, imm12 int32) {
	bytes := castToBytes(reg[rs2])
	Mem[reg[rs1]+imm12] = bytes[0]
	Mem[reg[rs1]+imm12+1] = bytes[1]
	Mem[reg[rs1]+imm12+2] = bytes[2]
	Mem[reg[rs1]+imm12+3] = bytes[3]
}

func sh(rs1 int32, rs2 int32, imm12 int32) {
	bytes := castToBytes(reg[rs2])
	Mem[reg[rs1]+imm12] = bytes[2]
	Mem[reg[rs1]+imm12+1] = bytes[3]
}

func sb(rs1 int32, rs2 int32, imm12 int32) {
	bytes := castToBytes(reg[rs2])
	Mem[reg[rs1]+imm12] = bytes[3]
}
