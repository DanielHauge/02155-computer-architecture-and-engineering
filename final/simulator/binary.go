package simulator

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"unsafe"
)

func readAsInt(b []byte) int32 {
	ui := binary.BigEndian.Uint32(b)
	return castToInt(&ui)
}

func readAsHalfWord(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func readAsUint(b []byte) uint32 {
	return uint32(big.NewInt(0).SetBytes(b).Uint64())
}

func castToBytes(i int32) []byte {
	bs := new(bytes.Buffer)
	binary.Write(bs, binary.LittleEndian, int32(i))
	return bs.Bytes()
}

func castUToBytes(i uint32) []byte {
	bs := new(bytes.Buffer)
	binary.Write(bs, binary.LittleEndian, uint32(i))
	return bs.Bytes()
}

func castToUint(i *int32) uint32 {
	return *(*uint32)(unsafe.Pointer(i))
}

func castToUint2(i int32) uint32 {
	return *(*uint32)(unsafe.Pointer(&i))
}

func castToInt(i *uint32) int32 {
	return *(*int32)(unsafe.Pointer(i))
}
