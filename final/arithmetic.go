package main

func add(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] + reg[rs2]
}

func sub(rd int, rs1 int, rs2 int) {
	reg[rd] = reg[rs1] - reg[rs2]
}
