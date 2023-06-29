package str

import "testing"

var Rune Gstr

func TestIsNumOrLetter(t *testing.T) {
	// Rune := Gstr{}
	if !Rune.IsNumOrLetter(rune('0')) {
		t.Fail()
	}
	if Rune.IsNumOrLetter(rune('@')) {
		t.Fail()
	}
}

func TestIsLetter(t *testing.T) {
	// Rune := Gstr{}
	if !Rune.IsLetter(rune('a')) {
		t.Fail()
	}
	if Rune.IsLetter(rune('0')) {
		t.Fail()
	}
}

func TestContainsChinese(t *testing.T) {
	if !Rune.ContainChinese("Hello 世界") {
		t.Fail()
	}
	if Rune.ContainChinese("Hello World") {
		t.Fail()
	}
	if Rune.ContainChinese("Hello，World") {
		t.Fail()
	}
}
