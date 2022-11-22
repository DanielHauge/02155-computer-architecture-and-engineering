package simulator

import (
	"testing"
)

func Test_readAsIntMinus(t *testing.T) {
	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(0xFF), byte(0xFF), byte(0xFF), byte(0xFF)
	i := readAsInt(bs)
	assert(i, -1, t)
}

func Test_readAsInt(t *testing.T) {
	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(0), byte(0), byte(0), byte(13)
	i := readAsInt(bs)
	assert(i, 13, t)
}

func Test_readAsUint(t *testing.T) {
	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(0), byte(0), byte(0), byte(13)
	i := readAsUint(bs)
	assert(i, 13, t)

}

func Test_readAsUintMaxNotMinus(t *testing.T) {
	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(0xFF), byte(0xFF), byte(0xFF), byte(0xFF)
	i := readAsUint(bs)
	assert(i, 4294967295, t)
}

func Test_castToBytesLength(t *testing.T) {
	assert(len(castToBytes(13)), 4, t)
}

func Test_castToBytes(t *testing.T) {
	bs := castToBytes(13)
	assertAll(bs[2:3], byte(0), t)
	assert(bs[0], byte(13), t)
}

func Test_castToUint(t *testing.T) {
	i := int32(-1)
	ui := castToUint(&i)
	assert(ui, 4294967295, t)

}
