package gtool

import "testing"

func TestBase64(t *testing.T) {
	str := Base64.EncodeToString([]byte("hahahah"))
	t.Log(str)
}
