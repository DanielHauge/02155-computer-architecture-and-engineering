package simulator

import "testing"

func TestMain(m *testing.M) {
	mem := make([]byte, 1000)
	Initialize(mem)
	m.Run()
}
