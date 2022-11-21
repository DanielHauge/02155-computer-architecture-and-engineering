package simulator

import "testing"

func Test_Decode(t *testing.T) {
	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(255), byte(255), byte(255), byte(255)
	decoded := Decode(bs)
	assert(decoded.opcode, 127, t)
	assert(decoded.rd, 31, t)
	assert(decoded.rs1, 31, t)
	assert(decoded.rs2, 31, t)
	assert(decoded.imm12_I, 4095, t)
	assert(decoded.imm12_SB, 4095, t)
	assert(decoded.imm20, 1048575, t)
	assert(decoded.funct3, 7, t)
	assert(decoded.funct7, 127, t)
}

func Test_R(t *testing.T) {
	testFunc := func(a int32, b int32, c int32) {
		assert(a, 25, t)
		assert(b, 52, t)
		assert(c, 32, t)
	}
	rFormat(instruction{rd: 25, rs1: 52, rs2: 32}, testFunc)
}

func Test_I(t *testing.T) {
	testFunc := func(a int32, b int32, c int32) {
		assert(a, 31, t)
		assert(b, 22, t)
		assert(c, 12, t)
	}
	iFormat(instruction{rd: 31, rs1: 22, imm12_I: 12}, testFunc)
}

func Test_U(t *testing.T) {
	testFunc := func(a int32, b int32) {
		assert(a, 31, t)
		assert(b, 252522, t)
	}
	uFormat(instruction{rd: 31, imm20: 252522}, testFunc)
}

func Test_S(t *testing.T) {
	testFunc := func(a int32, b int32, c int32) {
		assert(a, 31, t)
		assert(b, 31, t)
		assert(c, 3131, t)
	}
	sFormat(instruction{rs1: 31, rs2: 31, imm12_SB: 3131}, testFunc)
}
