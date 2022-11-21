package simulator

import (
	"testing"
)

func TestMain(m *testing.M) {
	mem := make([]byte, 1000)
	Initialize(mem)
	m.Run()
}

type Comparable interface {
	int32 | byte | uint32 | int
}

func assert[T Comparable](actual T, expected T, t *testing.T) {
	if expected != actual {
		t.Log("Expected ", expected, " But got ", actual)
		t.Fail()
	}
}

func assertAll[T Comparable](actuals []T, expected T, t *testing.T) {
	for _, actual := range actuals {
		assert(actual, expected, t)
	}
}
