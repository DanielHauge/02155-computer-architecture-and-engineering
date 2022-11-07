package main

func beq(rs1 int, rs2 int, imm12 int) {
	if rs1 == rs2 {
		pc += imm12
	}
}
