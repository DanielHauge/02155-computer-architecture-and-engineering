package simulator

import "testing"

func Test_add(t *testing.T) {
	reg[5] = 5
	reg[6] = 3
	add(1, 5, 6)
	assert(reg[1], 8, t)
}

func Test_sub(t *testing.T) {
	reg[5] = 5
	reg[6] = 3
	sub(1, 5, 6)
	assert(reg[1], 2, t)
}

func Test_addi(t *testing.T) {
	reg[5] = 5
	addi(1, 5, 6)
	assert(reg[1], 11, t)
}

func Test_slt(t *testing.T) {
	reg[5] = 6
	reg[6] = 5
	slt(1, 5, 6)
	assert(reg[1], 0, t)
}

func Test_slti(t *testing.T) {
	reg[5] = 5
	slti(1, 5, 6)
	assert(reg[1], 1, t)
}

func Test_sltu(t *testing.T) {
	reg[5] = -1
	reg[6] = 1
	sltu(1, 5, 6)
	assert(reg[1], 0, t)
}

func Test_sltiu(t *testing.T) {
	reg[5] = 1
	sltiu(1, 5, -1)
	assert(reg[1], 1, t)
}

func Test_lui(t *testing.T) {
	lui(1, 1)
	assert(reg[1], 4096, t)
}

func Test_auip(t *testing.T) {
	Pc = 12
	auip(1, 1)
	assert(reg[1], 4108, t)
}
