#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail
if [[ "${TRACE-0}" == "1" ]]; then
    set -o xtrace
fi

if [[ "${1-}" =~ ^-*h(elp)?$ ]]; then
    echo 'Usage: ./compile_tests.sh <filename>

This is an awesome bash script to make your life better. The filename should not include .s and should have a .s file

'
    exit
fi

cd "$(dirname "$0")"

main() {
    riscv64-linux-gnu-gcc -march=rv32i -mabi=ilp32 -c assembly/$1.s -o object/$1.o
    riscv64-linux-gnu-objcopy -O binary object/$1.o binary/$1.bin
}

main "$@"
