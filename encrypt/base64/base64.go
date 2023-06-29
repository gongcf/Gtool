package base64

import (
	"encoding/base64"
	"io/ioutil"
)

type GBase64 struct{}

// Encode encodes bytes with BASE64 algorithm.
func (*GBase64) Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// EncodeString encodes string with BASE64 algorithm.
func (gb *GBase64) EncodeString(src string) string {
	return gb.EncodeToString([]byte(src))
}

// EncodeToString encodes bytes to string with BASE64 algorithm.
func (gb *GBase64) EncodeToString(src []byte) string {
	return string(gb.Encode(src))
}

// EncryptFile encodes file content of <path> using BASE64 algorithms.
func (gb *GBase64) EncodeFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return gb.Encode(content), nil
}

// MustEncodeFile encodes file content of <path> using BASE64 algorithms.
// It panics if any error occurs.
func (gb *GBase64) MustEncodeFile(path string) []byte {
	result, err := gb.EncodeFile(path)
	if err != nil {
		panic(err)
	}
	return result
}

// EncodeFileToString encodes file content of <path> to string using BASE64 algorithms.
func (gb *GBase64) EncodeFileToString(path string) (string, error) {
	content, err := gb.EncodeFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// MustEncodeFileToString encodes file content of <path> to string using BASE64 algorithms.
// It panics if any error occurs.
func (gb *GBase64) MustEncodeFileToString(path string) string {
	result, err := gb.EncodeFileToString(path)
	if err != nil {
		panic(err)
	}
	return result
}

// Decode decodes bytes with BASE64 algorithm.
func (*GBase64) Decode(data []byte) ([]byte, error) {
	src := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(src, data)
	return src[:n], err
}

// MustDecode decodes bytes with BASE64 algorithm.
// It panics if any error occurs.
func (gb *GBase64) MustDecode(data []byte) []byte {
	result, err := gb.Decode(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeString decodes string with BASE64 algorithm.
func (gb *GBase64) DecodeString(data string) ([]byte, error) {
	return gb.Decode([]byte(data))
}

// MustDecodeString decodes string with BASE64 algorithm.
// It panics if any error occurs.
func (gb *GBase64) MustDecodeString(data string) []byte {
	result, err := gb.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeString decodes string with BASE64 algorithm.
func (gb *GBase64) DecodeToString(data string) (string, error) {
	b, err := gb.DecodeString(data)
	return string(b), err
}

// MustDecodeToString decodes string with BASE64 algorithm.
// It panics if any error occurs.
func (gb *GBase64) MustDecodeToString(data string) string {
	result, err := gb.DecodeToString(data)
	if err != nil {
		panic(err)
	}
	return result
}
