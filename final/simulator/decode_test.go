package simulator

import "testing"

func Test_Decode(t *testing.T) {
	var instr uint32 = 0b10000000000000000000010100110111
	bs := castUToBytes(instr)
	decoded := Decode(bs)
	assert(decoded.opcode, 0b0110111, t)
}

func Test_Decode_S(t *testing.T) {
	var instr uint32 = 0b1_000000_0000000000000_01010_0110111
	bs := castUToBytes(instr)
	decoded := Decode(bs)
	assert(uint32(decoded.Imm_S), 0b111111111111111111111_000000_01010, t)
}

func Test_Decode_B(t *testing.T) {
	var instr uint32 = 0b1_111111_0101101010001_1110_1_1100011
	bs := castUToBytes(instr)
	decoded := Decode(bs)
	assert(decoded.Imm_B, 0b1111111111111111111_11_111111_1110_0, t)
}

func Test_Decode_U(t *testing.T) {
	var instr uint32 = 0b00010010001101000101001010110111
	bs := castUToBytes(instr)
	decoded := Decode(bs)
	assert(decoded.Imm_U, 0x12345, t)
}

func Test_R(t *testing.T) {
	testFunc := func(a uint32, b uint32, c uint32) {
		assert(a, 25, t)
		assert(b, 52, t)
		assert(c, 32, t)
	}
	rFormat(instruction{rd: 25, rs1: 52, rs2: 32}, testFunc)
}

func Test_I(t *testing.T) {
	testFunc := func(a uint32, b uint32, c uint32) {
		assert(a, 31, t)
		assert(b, 22, t)
		assert(c, 12, t)
	}
	iFormat(instruction{rd: 31, rs1: 22, Imm_I: 12}, testFunc)
}

func Test_U(t *testing.T) {
	testFunc := func(a uint32, b uint32) {
		assert(a, 31, t)
		assert(b, 252522, t)
	}
	uFormat(instruction{rd: 31, Imm_U: 252522}, testFunc)
}

func Test_J(t *testing.T) {
	testFunc := func(a uint32, b uint32) {
		assert(a, 31, t)
		assert(b, 252522, t)
	}
	jFormat(instruction{rd: 31, Imm_J: 252522}, testFunc)
}

func Test_S(t *testing.T) {
	testFunc := func(a uint32, b uint32, c uint32) {
		assert(a, 31, t)
		assert(b, 31, t)
		assert(c, 3131, t)
	}
	sFormat(instruction{rs1: 31, rs2: 31, Imm_S: 3131}, testFunc)
}

func Test_B(t *testing.T) {
	testFunc := func(a uint32, b uint32, c uint32) {
		assert(a, 31, t)
		assert(b, 31, t)
		assert(c, 3131, t)
	}
	bFormat(instruction{rs1: 31, rs2: 31, Imm_B: 3131}, testFunc)
}
