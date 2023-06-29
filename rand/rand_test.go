package rand

import "testing"

var rd GRand

func TestInts(t *testing.T) {
	ints := rd.Ints(10, 19, 20)
	if 9 != len(ints) {
		t.Fail()
	}
	ints = rd.Ints(10, 19, 5)
	if 5 != len(ints) {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	r1 := rd.String(16)
	r2 := rd.String(16)

	if r1 == r2 {
		t.Fail()
	}
}

func TestInt(t *testing.T) {
	r1 := rd.Int(0, 65535)
	r2 := rd.Int(0, 65535)

	if r1 == r2 {
		t.Fail()
	}
}
