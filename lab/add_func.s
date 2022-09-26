
.text
main:
addi a1, x0, 5
addi a2, x0, 6
jal ra, add

addi a1, a0, 0
addi a0, x0, 1
ecall # print pesult

addi a0, x0, 10
ecall # exit

add:
add a0, a1, a2
jalr x0,0(x1) # return

