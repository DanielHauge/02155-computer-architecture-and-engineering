package simulator

import (
	"testing"
)

func Test_readAsIntMinus(t *testing.T) {

	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(0xFF), byte(0xFF), byte(0xFF), byte(0xFF)
	i := readAsInt(bs)
	if i != -1 {
		t.Log("Expected -1, but got", i)
		t.Fail()
	}
}

func Test_readAsInt(t *testing.T) {

	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(0), byte(0), byte(0), byte(13)
	i := readAsInt(bs)
	if i != 13 {
		t.Log("Expected 13, but got", i)
		t.Fail()
	}
}

func Test_readAsUint(t *testing.T) {
	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(0), byte(0), byte(0), byte(13)
	i := readAsUint(bs)
	if i != 13 {
		t.Log("Expected 13, but got ", i)
		t.Fail()
	}
}

func Test_readAsUintMaxNotMinus(t *testing.T) {
	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(0xFF), byte(0xFF), byte(0xFF), byte(0xFF)
	i := readAsUint(bs)
	if i != 4294967295 {
		t.Log("Expected 4294967295, but got ", i)
		t.Fail()
	}
}

func Test_castToBytesLength(t *testing.T) {
	if len(castToBytes(13)) != 4 {
		t.Log("Expected length 4, but got ", len(castToBytes(13)))
		t.Fail()
	}
}

func Test_castToBytes(t *testing.T) {
	bs := castToBytes(13)
	if bs[0] != byte(0) && bs[1] != byte(0) && bs[2] != byte(0) && bs[3] != byte(13) {
		t.Log("Expected bytes 0:0:0:13 , but got ", bs)
		t.Fail()
	}
}

func Test_castToUint(t *testing.T) {
	i := int32(-1)
	ui := castToUint(&i)
	if ui != 4294967295 {
		t.Log("Expected bytes 4294967295, but got ", ui)
		t.Fail()
	}

}
