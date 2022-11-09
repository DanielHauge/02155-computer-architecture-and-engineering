package main

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

var (
	endian = binary.BigEndian
)

func readAsInt(b []byte) int {
	var theInt int
	err := binary.Read(bytes.NewReader(b), endian, &theInt)
	if err != nil {
		panic(err)
	}
	return theInt
}

func readAsUint(b []byte) uint {
	var theInt uint
	err := binary.Read(bytes.NewReader(b), endian, &theInt)
	if err != nil {
		panic(err)
	}
	return theInt
}

func castToBytes(i int) []byte {
	bs := new(bytes.Buffer)
	binary.Write(bs, endian, int32(i))
	return bs.Bytes()
}

func castToUint(i *int) uint {
	return *(*uint)(unsafe.Pointer(i))
}

// func castToInt(i *uint) int {
// 	return *(*int)(unsafe.Pointer(i))
// }
