package file

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

var testdataDir = "testdata"
var gf GFile

func TestRemoveEmptyDirs(t *testing.T) {
	testPath := "testdata/dir"

	// case 1
	if err := os.RemoveAll(testPath); nil != err {
		t.Errorf("clear test empty dir [%s] failed: %s", testPath, err)
	}

	a := filepath.Join(testPath, "a")
	if err := os.MkdirAll(a, 0755); nil != err {
		t.Errorf("make dir [%s] failed: %s", testPath, err)
	}

	if err := gf.RemoveEmptyDirs(testPath); nil != err {
		t.Errorf("remove empty dirs failed: %s", err)
	}

	if gf.IsDir(a) || gf.IsDir(testPath) {
		t.Errorf("empty dir [%s] exists", a)
	}

	// case 2
	if err := os.RemoveAll(testPath); nil != err {
		t.Errorf("clear test empty dir [%s] failed: %s", testPath, err)
	}

	if err := os.MkdirAll(a, 0755); nil != err {
		t.Errorf("make dir [%s] failed: %s", testPath, err)
	}
	test := filepath.Join(a, "test")
	if err := os.WriteFile(test, []byte(""), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", test, err)
	}

	if err := gf.RemoveEmptyDirs(testPath); nil != err {
		t.Errorf("remove empty dirs failed: %s", err)
	}

	if !gf.IsDir(a) || !gf.IsDir(testPath) {
		t.Errorf("empty dir [%s] exists", a)
	}

	// case 3
	if err := os.RemoveAll(testPath); nil != err {
		t.Errorf("clear test empty dir [%s] failed: %s", testPath, err)
	}

	if err := os.MkdirAll(a, 0755); nil != err {
		t.Errorf("make dir [%s] failed: %s", testPath, err)
	}

	if err := gf.RemoveEmptyDirs(testPath, "a"); nil != err {
		t.Errorf("remove empty dirs failed: %s", err)
	}

	if !gf.IsDir(a) || !gf.IsDir(testPath) {
		t.Errorf("empty dir [%s] exists", a)
	}

	if err := os.RemoveAll(testPath); nil != err {
		t.Errorf("clear test empty dir [%s] failed: %s", testPath, err)
	}
}

func TestIsValidFilename(t *testing.T) {
	if !gf.IsValidFilename("hello.go") {
		t.Errorf("[hello.go] should be a valid filename")
	}
	if gf.IsValidFilename("hello?.go") {
		t.Errorf("[hello?.go] should not be a valid filename")
	}
}

func TestWriteFileSaferByReader(t *testing.T) {
	writePath := "testdata/filewrite.go"
	defer os.RemoveAll(writePath)
	if err := gf.WriteFileSaferByReader(writePath, strings.NewReader("test"), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", writePath, err)
	}
}

func TestWriteFileSaferWithoutChangeTime(t *testing.T) {
	writePath := "testdata/filewrite.go"
	defer os.RemoveAll(writePath)

	if err := os.WriteFile(writePath, []byte("0"), 0644); nil != err {
		t.Fatalf("write file [%s] failed: %s", writePath, err)
	}

	info, err := os.Stat(writePath)
	if nil != err {
		t.Fatalf("stat file [%s] failed: %s", writePath, err)
	}
	modTime1 := info.ModTime()

	if err = gf.WriteFileSaferWithoutChangeTime(writePath, []byte("test"), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", writePath, err)
	}

	info, err = os.Stat(writePath)
	if nil != err {
		t.Fatalf("stat file [%s] failed: %s", writePath, err)
	}
	modTime2 := info.ModTime()
	if modTime1 != modTime2 {
		t.Errorf("mod time should not be changed")
	}

	writePath1 := "testdata/filewrite1.go"
	defer os.RemoveAll(writePath1)
	if err = gf.WriteFileSaferWithoutChangeTime(writePath1, []byte("test"), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", writePath, err)
	}
	info, err = os.Stat(writePath1)
	if nil != err {
		t.Fatalf("stat file [%s] failed: %s", writePath1, err)
	}
	modTime3 := info.ModTime()
	t.Logf("file mod time [%v]", modTime3)
}

func TestWriteFileSafer(t *testing.T) {
	writePath := "testdata/filewrite.go"
	defer os.RemoveAll(writePath)

	if err := os.WriteFile(writePath, []byte("0"), 0644); nil != err {
		t.Fatalf("write file [%s] failed: %s", writePath, err)
	}

	info, err := os.Stat(writePath)
	if nil != err {
		t.Fatalf("stat file [%s] failed: %s", writePath, err)
	}

	if err = gf.WriteFileSafer(writePath, []byte("test"), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", writePath, err)
	}

	info, err = os.Stat(writePath)
	if nil != err {
		t.Fatalf("stat file [%s] failed: %s", writePath, err)
	}
	modTime2 := info.ModTime()
	t.Logf("file mod time [%v]", modTime2)
}

func TestIsHidden(t *testing.T) {
	filename := "./file.go"
	isHidden := gf.IsHidden(filename)
	if isHidden {
		t.Error("file [" + filename + "] is not hidden")
	}
}

func TestGetFileSize(t *testing.T) {
	filename := "./file.go"
	size := gf.GetFileSize(filename)
	if 0 > size {
		t.Error("file [" + filename + "] size is [" + strconv.FormatInt(size, 10) + "]")
	}
}

func TestIsExist(t *testing.T) {
	if !gf.IsExist(".") {
		t.Error(". must exist")
		return
	}
}

func TestIdBinary(t *testing.T) {
	if gf.IsBinary("not binary content") {
		t.Error("The content should not be binary")
		return
	}
}

func TestIsImg(t *testing.T) {
	if !gf.IsImg(".jpg") {
		t.Error(".jpg should be a valid extension of a image file")
		return
	}
}

func TestIsDir(t *testing.T) {
	if !gf.IsDir(".") {
		t.Error(". should be a directory")
		return
	}
}

func TestCopyDir(t *testing.T) {
	source := "testcopydir"
	os.Mkdir(source, 0644)
	dest := filepath.Join(testdataDir, source)
	defer os.Remove(dest)

	err := gf.CopyDir(source, dest)
	if nil != err {
		t.Error("Copy dir failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}
}

func TestCopyFile(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := gf.CopyFile(source, dest)
	if nil != err {
		t.Error("Copy file failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}
}

func TestCopy(t *testing.T) {
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := gf.Copy("./file.go", dest)
	if nil != err {
		t.Error("Copy failed: ", err)
		return
	}

	sourceStat, _ := os.Stat("./file.go")
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}
}

func TestCopyDirNewtimes(t *testing.T) {
	source := "testcopydir"
	os.Mkdir(source, 0644)
	dest := filepath.Join(testdataDir, source)
	defer os.Remove(dest)

	time.Sleep(100 * time.Millisecond) // CI

	err := gf.CopyDirNewtimes(source, dest)
	if nil != err {
		t.Error("Copy dir failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}

func TestCopyFileNewtimes(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := gf.CopyFileNewtimes(source, dest)
	if nil != err {
		t.Error("Copy file failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}

func TestCopyNewtimes(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := gf.CopyNewtimes(source, dest)
	if nil != err {
		t.Error("Copy failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}
