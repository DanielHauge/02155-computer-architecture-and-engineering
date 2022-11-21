package simulator

import "testing"

func Test_and(t *testing.T) {
	reg[5] = 15
	reg[6] = 8
	and(1, 5, 6)
	assert(reg[1], 8, t)
}

func Test_or(t *testing.T) {
	reg[5] = 7
	reg[6] = 8
	or(1, 5, 6)
	assert(reg[1], 15, t)
}

func Test_xor(t *testing.T) {
	reg[5] = 3
	reg[6] = 15
	xor(1, 5, 6)
	assert(reg[1], 12, t)
}

func Test_andi(t *testing.T) {
	reg[5] = 8
	andi(1, 5, 15)
	assert(reg[1], 8, t)
}

func Test_ori(t *testing.T) {
	reg[5] = 7
	ori(1, 5, 8)
	assert(reg[1], 15, t)
}

func Test_xori(t *testing.T) {
	reg[5] = 3
	xori(1, 5, 15)
	assert(reg[1], 12, t)
}

func Test_sll(t *testing.T) {
	reg[5] = 2
	reg[6] = 3
	sll(1, 5, 6)
	assert(reg[1], 16, t)
}

func Test_srl(t *testing.T) {
	reg[5] = 16
	reg[6] = 3
	srl(1, 5, 6)
	assert(reg[1], 2, t)
}

func Test_sra(t *testing.T) {
	reg[5] = -1
	reg[6] = 2
	sra(1, 5, 6)
	assert(reg[1], -1, t)
}

func Test_slli(t *testing.T) {
	reg[5] = 2
	slli(1, 5, 3)
	assert(reg[1], 16, t)
}

func Test_srli(t *testing.T) {
	reg[5] = 16
	srli(1, 5, 3)
	assert(reg[1], 2, t)
}

func Test_srai(t *testing.T) {
	reg[5] = -1
	srai(1, 5, 2)
	assert(reg[1], -1, t)
}
