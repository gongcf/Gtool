package str

import (
	"testing"
)

var gstr Gstr

func TestEqual(t *testing.T) {
	if !gstr.Equal([]string{"foo", "bar"}, []string{"foo", "bar"}) {
		t.Error("[\"foo\", \"bar\"] and [\"foo\", \"bar\"] should be equal")
		return
	}

	if gstr.Equal([]string{"foo", "bar"}, []string{"foo", "bar", "baz"}) {
		t.Error("[\"foo\", \"bar\"] and [\"foo\", \"bar\", \"baz\"] should not be equal")
		return
	}

	if gstr.Equal([]string{"foo", "bar"}, []string{"bar", "foo"}) {
		t.Error("[\"foo\", \"bar\"] and [\"bar\", \"foo\"] should not be equal")
		return
	}
}

func TestRemoveDuplicatedElem(t *testing.T) {
	expected := []string{"foo", "bar", "baz"}
	got := gstr.RemoveDuplicatedElem([]string{"foo", "bar", "bar", "baz", "baz"})
	if !gstr.Equal(expected, got) {
		t.Errorf("expected is %v, but got is %v", expected, got)
	}
}

func TestExcludeElem(t *testing.T) {
	got := gstr.ExcludeElem([]string{"foo", "bar", "baz"}, []string{"bar", "baz"})
	if 1 != len(got) || "foo" != got[0] {
		t.Errorf("expected [foo], got [%s]", got)
		return
	}
}

func TestRemoveElem(t *testing.T) {
	got := gstr.RemoveElem([]string{"foo", "bar", "baz"}, "bar")
	if 2 != len(got) || "foo" != got[0] || "baz" != got[1] {
		t.Errorf("expected [foo, baz], got [%s]", got)
		return
	}
}

func TestSubstringsBetween(t *testing.T) {
	got := gstr.SubstringsBetween("foo<bar>baz<bar2>", "<", ">")
	if 2 != len(got) {
		t.Errorf("substrings between [%s] should have 2 elements", got)
		return
	}
}

func TestIsASCII(t *testing.T) {
	if !gstr.IsASCII("foo") {
		t.Error("[foo] should be ASCII")
		return
	}

	if gstr.IsASCII("foo测试") {
		t.Error("[foo测试] should not be ASCII")
		return
	}
}

func TestRemoveInvisible(t *testing.T) {
	expected := "foo测试barbaz"
	got := gstr.RemoveInvisible("foo\u200b测试\nbar\tbaz")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}

func TestToBytes(t *testing.T) {
	str := "Gulu 你好！"
	bytes := gstr.ToBytes(str)
	if str2 := gstr.FromBytes(bytes); str != str2 {
		t.Errorf("Str Bytes convert failed [str=%s, str2=%s]", str, str2)
	}
}

func TestContains(t *testing.T) {
	if !gstr.Contains("123", []string{"123", "345"}) {
		t.Error("[\"123\", \"345\"] should contain \"123\"")
		return
	}
}

func TestReplaceIgnoreCase(t *testing.T) {
	expected := "Foabcdr"
	got := gstr.ReplaceIgnoreCase("Foobar", "oBa", "abcd")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}

func TestReplacesIgnoreCase(t *testing.T) {
	expected := "abcdbarefgh"
	got := gstr.ReplacesIgnoreCase("Foobarbaz", "foo", "abcd", "baz", "efgh")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}

	expected = "bar baz baz"
	got = gstr.ReplacesIgnoreCase("foo bar baz", "foo", "bar", "bar", "baz")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}

	expected = "bar baz baz"
	got = gstr.ReplacesIgnoreCase("foo bar baz", "Bar", "baz", "foo", "bar")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}

	expected = "fazz bar barr"
	got = gstr.ReplacesIgnoreCase("foo bar baz", "oo", "azz", "az", "arr")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}

func TestEncloseIgnoreCase(t *testing.T) {
	var expected, got string
	expected = "<mark>Foo</mark>bar<mark>baz</mark>"
	got = gstr.EncloseIgnoreCase("Foobarbaz", "<mark>", "</mark>", "foo", "baz")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}

	expected = "F<mark>oo</mark><mark>ba</mark>r<mark>ba</mark>z"
	got = gstr.EncloseIgnoreCase("Foobarbaz", "<mark>", "</mark>", "Oo", "Ba")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}

func TestLCS(t *testing.T) {
	str := gstr.LCS("123456", "abc34def")
	if "34" != str {
		t.Error("[\"123456\"] and [\"abc34def\"] should have the longest common substring [\"34\"]")
		return
	}
}

func TestSubStr(t *testing.T) {
	expected := "foo测"
	got := gstr.SubStr("foo测试bar", 4)
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}
