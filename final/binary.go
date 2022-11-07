package main

import (
	"bytes"
	"encoding/binary"
)

var (
	endian = binary.BigEndian
)

func castToInt(b []byte) int {
	var theInt int
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
