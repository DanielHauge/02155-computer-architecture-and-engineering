
.text
main:
addi a1, x0, 8
jal ra, factorial

addi a1, a0, 0
addi a0, x0, 1
ecall # print pesult
beq x0, x0 End

factorial:

BNEZ a1, Else
addi a0, x0, 1
jalr x0,0(ra)
beq x0,x0,Exit

Else: 
addi a1 a1 -1
addi sp,sp,-4
sw ra 0(sp)
jal ra, factorial
addi a1 a1 1
mul a0 a1 a0
lw ra,0(sp)
addi sp,sp,4

Exit:
jalr x0,0(ra) # return

End:
