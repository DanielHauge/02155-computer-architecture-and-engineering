package simulator

import "testing"

func Test_Initialize(t *testing.T) {
	customMemory := make([]byte, 1000)
	customMemory[0] = byte(255)
	customMemory[5] = byte(255)
	customMemory[255] = byte(255)
	Mem[0] = byte(125)
	Pc = 25
	reg[2] = 25252
	Initialize(customMemory)
	assert(Mem[0], byte(255), t)
	assert(Mem[5], byte(255), t)
	assert(Mem[255], byte(255), t)
	assert(Pc, 0, t)
	assert(reg[2], 1000, t)
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
	reg[5] = 26
	lw(1, 5, 0)
	assert(reg[1], -1, t)
	lw(1, 5, 4)
	assert(reg[1], 1, t)
}

func Test_lh(t *testing.T) {
	Mem[26] = byte(255)
	Mem[27] = byte(255)
	Mem[28] = byte(255)
	Mem[29] = byte(255)
	Mem[30] = byte(0)
	Mem[31] = byte(0)
	Mem[32] = byte(0)
	Mem[33] = byte(1)
	reg[5] = 26
	lh(1, 5, 3)
	assert(reg[1], 65280, t)
	lh(1, 5, 6)
	assert(reg[1], 1, t)
}

func Test_lb(t *testing.T) {
	Mem[26] = byte(255)
	Mem[27] = byte(255)
	Mem[28] = byte(255)
	Mem[29] = byte(255)
	Mem[30] = byte(0)
	Mem[31] = byte(0)
	Mem[32] = byte(0)
	Mem[33] = byte(1)
	reg[5] = 26
	lb(1, 5, 3)
	assert(reg[1], 255, t)
	lb(1, 5, 6)
	assert(reg[1], 0, t)
	lb(1, 5, 7)
	assert(reg[1], 1, t)
}

func Test_lwu(t *testing.T) {
	Mem[26] = byte(255)
	Mem[27] = byte(255)
	Mem[28] = byte(255)
	Mem[29] = byte(255)
	reg[5] = 26
	lwu(1, 5, 0)
	assert(reg[1], -1, t)
}

func Test_lhu(t *testing.T) {
	Mem[26] = byte(255)
	Mem[27] = byte(255)
	reg[5] = 26
	lwu(1, 5, 0)
	assert(reg[1], -1, t)
}

func Test_lbu(t *testing.T) {
	Mem[26] = byte(255)
	reg[5] = 26
	lbu(1, 5, 0)
	assert(reg[1], 255, t)
}

func Test_sw(t *testing.T) {
	reg[5] = -1
	reg[6] = 62
	sw(6, 5, 2)
	assert(Mem[63], byte(0), t)
	assert(Mem[64], byte(255), t)
	assert(Mem[65], byte(255), t)
	assert(Mem[66], byte(255), t)
	assert(Mem[67], byte(255), t)
	assert(Mem[68], byte(0), t)
}

func Test_sh(t *testing.T) {
	reg[5] = -1
	reg[6] = 82
	sh(6, 5, 2)
	assert(Mem[83], byte(0), t)
	assert(Mem[84], byte(255), t)
	assert(Mem[85], byte(255), t)
	assert(Mem[86], byte(0), t)
}

func Test_sb(t *testing.T) {
	reg[5] = -1
	reg[6] = 22
	sb(6, 5, 2)
	assert(Mem[23], byte(0), t)
	assert(Mem[24], byte(255), t)
	assert(Mem[25], byte(0), t)
}
