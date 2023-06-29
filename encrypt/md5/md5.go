package md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

type GMd5 struct{}

// Encrypt encrypts any type of variable using MD5 algorithms.
// It uses gconv package to convert <v> to its bytes type.
func (gm *GMd5) Encrypt(data []byte) (encrypt string, err error) {
	return gm.EncryptBytes(data)
}

// MustEncrypt encrypts any type of variable using MD5 algorithms.
// It uses gconv package to convert <v> to its bytes type.
// It panics if any error occurs.
func (gm *GMd5) MustEncrypt(data []byte) string {
	result, err := gm.Encrypt(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptBytes encrypts <data> using MD5 algorithms.
func (*GMd5) EncryptBytes(data []byte) (encrypt string, err error) {
	h := md5.New()
	if _, err = h.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptBytes encrypts <data> using MD5 algorithms.
// It panics if any error occurs.
func (gm *GMd5) MustEncryptBytes(data []byte) string {
	result, err := gm.EncryptBytes(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptBytes encrypts string <data> using MD5 algorithms.
func (gm *GMd5) EncryptString(data string) (encrypt string, err error) {
	return gm.EncryptBytes([]byte(data))
}

// MustEncryptString encrypts string <data> using MD5 algorithms.
// It panics if any error occurs.
func (gm *GMd5) MustEncryptString(data string) string {
	result, err := gm.EncryptString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptFile encrypts file content of <path> using MD5 algorithms.
func (*GMd5) EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptFile encrypts file content of <path> using MD5 algorithms.
// It panics if any error occurs.
func (gm *GMd5) MustEncryptFile(path string) string {
	result, err := gm.EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
