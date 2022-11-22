package simulator

import "testing"

func Test_Initialize(t *testing.T) {
	customMemory := make([]byte, 1000)
	customMemory[0] = byte(255)
	customMemory[5] = byte(255)
	customMemory[255] = byte(255)
	Mem[0] = byte(125)
	Pc = 25
	Reg[2] = 25252
	Initialize(customMemory)
	assert(Mem[0], byte(255), t)
	assert(Mem[5], byte(255), t)
	assert(Mem[255], byte(255), t)
	assert(Pc, 0, t)
	assert(Reg[2], 1000, t)
}

func Test_lw(t *testing.T) {
	Mem[26] = byte(255)
	Mem[27] = byte(255)
	Mem[28] = byte(255)
	Mem[29] = byte(255)
	Mem[30] = byte(0)
	Mem[31] = byte(0)
	Mem[32] = byte(0)
	Mem[33] = byte(1)
	Reg[5] = 26
	lw(1, 5, 0)
	assert(Reg[1], castToUint2(-1), t)
	lw(1, 5, 4)
	assert(Reg[1], 1, t)
}

func Test_lwu(t *testing.T) {
	Mem[26] = byte(255)
	Mem[27] = byte(255)
	Mem[28] = byte(255)
	Mem[29] = byte(255)
	Reg[5] = 26
	lwu(1, 5, 0)
	assert(Reg[1], castToUint2(-1), t)
}

func Test_lhu(t *testing.T) {
	Mem[26] = byte(255)
	Mem[27] = byte(255)
	Reg[5] = 26
	lwu(1, 5, 0)
	assert(Reg[1], castToUint2(-1), t)
}

func Test_lbu(t *testing.T) {
	Mem[26] = byte(255)
	Reg[5] = 26
	lbu(1, 5, 0)
	assert(Reg[1], 255, t)
}

func Test_sw(t *testing.T) {
	Reg[5] = castToUint2(-1)
	Reg[6] = 62
	sw(6, 5, 2)
	assert(Mem[63], byte(0), t)
	assert(Mem[64], byte(255), t)
	assert(Mem[65], byte(255), t)
	assert(Mem[66], byte(255), t)
	assert(Mem[67], byte(255), t)
	assert(Mem[68], byte(0), t)
}

func Test_sh(t *testing.T) {
	Reg[5] = castToUint2(-1)
	Reg[6] = 82
	sh(6, 5, 2)
	assert(Mem[83], byte(0), t)
	assert(Mem[84], byte(255), t)
	assert(Mem[85], byte(255), t)
	assert(Mem[86], byte(0), t)
}

func Test_sb(t *testing.T) {
	Reg[5] = castToUint2(-1)
	Reg[6] = 22
	sb(6, 5, 2)
	assert(Mem[23], byte(0), t)
	assert(Mem[24], byte(255), t)
	assert(Mem[25], byte(0), t)
}
