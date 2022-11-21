package simulator

import "testing"

func Test_beq(t *testing.T) {
	Pc = 12
	Reg[5] = 5
	Reg[6] = 5
	beq(5, 6, 12)
	assert(Pc, 24, t)
	Reg[6] = 6
	beq(5, 6, 12)
	assert(Pc, 24, t)
}

func Test_bne(t *testing.T) {
	Pc = 12
	Reg[5] = 5
	Reg[6] = 6
	bne(5, 6, 12)
	assert(Pc, 24, t)
	Reg[6] = 5
	bne(5, 6, 12)
	assert(Pc, 24, t)
}

func Test_bge(t *testing.T) {
	Pc = 12
	Reg[5] = 5
	Reg[6] = 6
	bge(5, 6, 12)
	assert(Pc, 12, t)
	Reg[5] = 6
	bge(5, 6, 12)
	assert(Pc, 24, t)
	Reg[5] = 12
	bge(5, 6, 2)
	assert(Pc, 26, t)
}

func Test_bgeu(t *testing.T) {
	Pc = 12
	Reg[5] = -1
	Reg[6] = 1
	bgeu(5, 6, 12)
	assert(Pc, 24, t)
}

func Test_blt(t *testing.T) {
	Pc = 12
	Reg[5] = 1
	Reg[6] = 1
	blt(5, 6, 12)
	assert(Pc, 12, t)
	Reg[5] = 0
	blt(5, 6, 12)
	assert(Pc, 24, t)
}

func Test_bltu(t *testing.T) {
	Pc = 12
	Reg[5] = -1
	Reg[6] = 1
	bltu(5, 6, 12)
	assert(Pc, 12, t)
	bltu(6, 5, 12)
	assert(Pc, 24, t)
}

func Test_jal(t *testing.T) {
	Pc = 12
	Reg[5] = 2
	jal(5, 20)
	assert(Reg[5], 16, t)
	assert(Pc, 32, t)
}

func Test_jalr(t *testing.T) {
	Pc = 14
	Reg[5] = 2
	Reg[6] = 4
	jalr(5, 6, 20)
	assert(Reg[5], 18, t)
	assert(Pc, 24, t)
}
