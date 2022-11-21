package simulator

var (
	Pc  int32
	Reg []int32
	Mem []byte
)

func Initialize(mem []byte) {
	Pc = 0
	Mem = mem
	Reg = make([]int32, 32)
	Reg[2] = int32(len(mem))
}

func lw(rd int32, rs1 int32, imm12 int32) {
	Reg[rd] = readAsInt(Mem[Reg[rs1]+imm12 : Reg[rs1]+imm12+4])
}

func lh(rd int32, rs1 int32, imm12 int32) {
	Reg[rd] = readAsInt(Mem[Reg[rs1]+imm12 : Reg[rs1]+imm12+2])
}

func lb(rd int32, rs1 int32, imm12 int32) {
	Reg[rd] = readAsInt([]byte{Mem[Reg[rs1]+imm12]})
}

func lwu(rd int32, rs1 int32, imm12 int32) {
	Reg[rd] = int32(readAsUint(Mem[Reg[rs1]+imm12 : Reg[rs1]+imm12+4]))
}

func lhu(rd int32, rs1 int32, imm12 int32) {
	Reg[rd] = int32(readAsUint(Mem[Reg[rs1]+imm12 : Reg[rs1]+imm12+2]))
}

func lbu(rd int32, rs1 int32, imm12 int32) {
	Reg[rd] = int32(readAsUint([]byte{Mem[Reg[rs1]+imm12]}))
}

func sw(rs1 int32, rs2 int32, imm12 int32) {
	bytes := castToBytes(Reg[rs2])
	Mem[Reg[rs1]+imm12] = bytes[0]
	Mem[Reg[rs1]+imm12+1] = bytes[1]
	Mem[Reg[rs1]+imm12+2] = bytes[2]
	Mem[Reg[rs1]+imm12+3] = bytes[3]
}

func sh(rs1 int32, rs2 int32, imm12 int32) {
	bytes := castToBytes(Reg[rs2])
	Mem[Reg[rs1]+imm12] = bytes[2]
	Mem[Reg[rs1]+imm12+1] = bytes[3]
}

func sb(rs1 int32, rs2 int32, imm12 int32) {
	bytes := castToBytes(Reg[rs2])
	Mem[Reg[rs1]+imm12] = bytes[3]
}
