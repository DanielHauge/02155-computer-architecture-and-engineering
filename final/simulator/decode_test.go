package simulator

import "testing"

func Test_Decode(t *testing.T) {
	bs := make([]byte, 4)
	bs[0], bs[1], bs[2], bs[3] = byte(123), byte(123), byte(123), byte(13)
	// // decoded := decode(bs)
	// if decoded.opcode != 13 {
	// 	t.Log("Opscode should be 32, it was ", decoded.opcode)
	// 	t.Fail()
	// }
	t.Fail()
}

func Test_R(t *testing.T) {
	t.Fail()
}

func Test_I(t *testing.T) {
	t.Fail()
}

func Test_U(t *testing.T) {
	t.Fail()
}

func Test_S(t *testing.T) {
	t.Fail()
}
