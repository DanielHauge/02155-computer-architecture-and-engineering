package simulator

import "testing"

func Test_add(t *testing.T) {
	reg[4] = 5
	reg[5] = 3
	add(1, 4, 5)
	if reg[1] != 8 {
		t.Log("Expected 8, but got", reg[1])
		t.Fail()
	}
}

func Test_sub(t *testing.T) {
	reg[4] = 5
	reg[5] = 3
	sub(1, 4, 5)
	if reg[1] != 2 {
		t.Log("Expected 2, but got", reg[1])
		t.Fail()
	}
}

func Test_addi(t *testing.T) {
	t.Fail()
}

func Test_slt(t *testing.T) {
	t.Fail()
}

func Test_slti(t *testing.T) {
	t.Fail()
}

func Test_sltu(t *testing.T) {
	t.Fail()
}

func Test_sltiu(t *testing.T) {
	t.Fail()
}

func Test_lui(t *testing.T) {
	t.Fail()
}

func Test_auip(t *testing.T) {
	t.Fail()
}
