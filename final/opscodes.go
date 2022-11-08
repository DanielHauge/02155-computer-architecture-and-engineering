package main

func wrap(op func(int, int, int), opFormatFunc func(instruction, func(int, int, int))) func(instruction) {
	return func(i instruction) { opFormatFunc(i, op) }
}

var operations = map[int]func(instruction){
	1000: wrap(add, R),
}
