package response

import "testing"

func TestNewResult(t *testing.T) {
	Ret := GResponse{}
	result := Ret.NewResult()
	if 0 != result.Code {
		t.Fail()
	}
}
