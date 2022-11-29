package simulator

import "fmt"

var Debug bool

type instruction struct {
	opcode uint32
	funct3 uint32
	funct7 uint32
	rd     uint32
	rs1    uint32
	rs2    uint32
	Imm_I  uint32
	Imm_S  uint32
	Imm_U  uint32
	Imm_B  uint32
	Imm_J  uint32
}

func Decode(instr_original []byte) instruction {

	instr := make([]byte, 4)
	copy(instr, instr_original)

	// Reverse cause of endian.
	for i, j := 0, len(instr)-1; i < j; i, j = i+1, j-1 {
		instr[i], instr[j] = instr[j], instr[i]
	}

	instrAsInt := readAsUint(instr)

	if Debug {
		fmt.Printf("%032b (%v):\t", instrAsInt, Pc)
	}

	signed := ((instrAsInt >> 31) & 0b1) == 1
	bit7b := ((instrAsInt >> 7) & 0b1) == 1
	var (
		Imm_IS_Mask, Imm_B_Mask, Imm_B_7Mask, Imm_J_Mask uint32
	)
	if signed {
		Imm_IS_Mask = 0b11111111111111111111100000000000
		Imm_B_Mask = 0b11111111111111111110000000000000
		Imm_J_Mask = 0b11111111111100000000000000000000
	} else {
		Imm_IS_Mask = 0
		Imm_B_Mask = 0
		Imm_J_Mask = 0
	}
	if bit7b {
		Imm_B_7Mask = 0b1100000000000 | Imm_B_Mask
	} else {
		Imm_B_7Mask = Imm_B_Mask
	}

	Imm_J := instrAsInt & 0b11111111111111111111000000000000
	Imm_J_12_19 := Imm_J & 0b11111111000000000000
	Imm_J_31 := (Imm_J & 0b100000000000000000000) >> 9
	Imm_21_30 := (Imm_J & 0b1111111111000000000000000000000) >> 20
	Imm_J_20 := (Imm_J & 0b100000000000000000000) >> 9

	Imm_J_final := Imm_J_12_19 | Imm_J_31 | Imm_21_30 | Imm_J_20 | Imm_J_Mask

	decoded := instruction{
		opcode: uint32(instrAsInt & 0b1111111),
		funct3: uint32((instrAsInt >> 12) & 0b111),
		funct7: uint32((instrAsInt >> 25) & 0b1111111),
		rd:     uint32((instrAsInt >> 7) & 0b11111),
		rs1:    uint32((instrAsInt >> 15) & 0b11111),
		rs2:    uint32((instrAsInt >> 20) & 0b11111),
		Imm_I:  uint32(((instrAsInt >> 20) & 0b111111111111) | Imm_IS_Mask),
		Imm_S:  uint32(((instrAsInt >> 20) & 0b11111100000) | ((instrAsInt >> 7) & 0b11111) | Imm_IS_Mask),
		Imm_U:  uint32(instrAsInt&0b11111111111111111111000000000000) >> 12,
		Imm_B:  uint32(((instrAsInt >> 20) & 0b11111100000) | ((instrAsInt >> 7) & 0b11110) | Imm_B_7Mask),
		Imm_J:  Imm_J_final,
	}

	return decoded
}

func bFormat(inst instruction, op func(uint32, uint32, uint32)) {
	if Debug {
		fmt.Printf("x%v x%v 0x%x\n", inst.rs1, inst.rs2, inst.Imm_B)
	}
	op(inst.rs1, inst.rs2, inst.Imm_B)
}

func rFormat(inst instruction, op func(uint32, uint32, uint32)) {
	if Debug {
		fmt.Printf("x%v x%v x%v\n", inst.rd, inst.rs1, inst.rs2)
	}
	if inst.rd == 0 {
		if Debug {
			fmt.Printf("Hov! Assigning x0 -> will ignore\n")
		}
		return
	}
	op(inst.rd, inst.rs1, inst.rs2)
}

func iFormat(inst instruction, op func(uint32, uint32, uint32)) {
	if Debug {
		fmt.Printf("x%v x%v 0x%x\n", inst.rd, inst.rs1, inst.Imm_I)
	}
	if inst.rd == 0 {
		if Debug {
			fmt.Printf("Hov! Assigning x0 -> will ignore\n")
		}
		return
	}
	op(inst.rd, inst.rs1, inst.Imm_I)
}

func uFormat(inst instruction, op func(uint32, uint32)) {
	if Debug {
		fmt.Printf("x%v 0x%x\n", inst.rd, inst.Imm_U)
	}
	if inst.rd == 0 {
		if Debug {
			fmt.Printf("Hov! Assigning x0 -> will ignore\n")
		}
		return
	}
	op(inst.rd, inst.Imm_U)
}

func jFormat(inst instruction, op func(uint32, uint32)) {
	if Debug {
		fmt.Printf("x%v 0x%v\n", inst.rd, inst.Imm_J)
	}
	if inst.rd == 0 {
		if Debug {
			fmt.Printf("Hov! Assigning x0 -> will ignore\n")
		}
		return
	}
	op(inst.rd, inst.Imm_J)
}

func sFormat(inst instruction, op func(uint32, uint32, uint32)) {
	if Debug {
		fmt.Printf("x%v x%v 0x%x\n", inst.rs1, inst.rs2, inst.Imm_S)
	}
	op(inst.rs1, inst.rs2, inst.Imm_S)
}
