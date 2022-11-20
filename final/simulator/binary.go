package simulator

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"unsafe"
)

var (
	endian = binary.BigEndian
)

func readAsInt(b []byte) int32 {
	return int32(big.NewInt(0).SetBytes(b).Int64())
}

func readAsUint(b []byte) uint32 {
	return uint32(big.NewInt(0).SetBytes(b).Uint64())
}

func castToBytes(i int32) []byte {
	bs := new(bytes.Buffer)
	binary.Write(bs, endian, int32(i))
	return bs.Bytes()
}

func castToUint(i *int32) uint32 {
	return *(*uint32)(unsafe.Pointer(i))
}
