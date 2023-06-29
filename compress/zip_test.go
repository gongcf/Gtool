package compress

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gongcf/gtool/file"
)

var testdataDir = "testdata"
var zipDirPath = filepath.Join(testdataDir, "test_zip")
var Zip GCompress
var File file.GFile

func TestCreate(t *testing.T) {
	zipFile, err := Zip.Create(zipDirPath + ".zip")
	if nil != err {
		t.Error(err)
		return
	}

	zipFile.AddDirectoryN(".", testdataDir)
	if nil != err {
		t.Error(err)
		return
	}

	err = zipFile.Close()
	if nil != err {
		t.Error(err)
		return
	}

	err = Zip.Unzip(zipDirPath+".zip", zipDirPath)
	if nil != err {
		t.Error(err)
		return
	}

	f1, _ := os.Stat(filepath.Join(testdataDir, "README.md"))
	f1ModTime := f1.ModTime()
	f2, _ := os.Stat(filepath.Join(zipDirPath, "README.md"))
	f2ModTime := f2.ModTime()
	if f1ModTime.Unix() != f2ModTime.Unix() {
		t.Error("ModTime error")
		return
	}
}

func TestUnzip(t *testing.T) {
	err := Zip.Unzip(zipDirPath+".zip", zipDirPath)
	if nil != err {
		t.Error(err)
		return
	}
}

func _TestEmptyDir(t *testing.T) {
	dir1 := "/dir/subDir1"
	dir2 := "/dir/subDir2"

	err := os.MkdirAll(zipDirPath+dir1, os.ModeDir)
	if nil != err {
		t.Error(err)

		return
	}

	err = os.MkdirAll(zipDirPath+dir2, os.ModeDir)
	if nil != err {
		t.Error(err)
		return
	}

	f, err := os.Create(zipDirPath + dir2 + "/file")
	if nil != err {
		t.Error(err)
		return
	}
	f.Close()

	zipFile, err := Zip.Create(zipDirPath + "/dir.zip")
	if nil != err {
		t.Error(err)
		return
	}

	zipFile.AddDirectoryN("dir", zipDirPath+"/dir")
	if nil != err {
		t.Error(err)
		return
	}

	err = zipFile.Close()
	if nil != err {
		t.Error(err)
		return
	}

	err = Zip.Unzip(zipDirPath+"/dir.zip", zipDirPath+"/unzipDir")
	if nil != err {
		t.Error(err)
		return
	}

	if !File.IsExist(zipDirPath+"/unzipDir") || !File.IsDir(zipDirPath+"/unzipDir") {
		t.Error("Unzip failed")
		return
	}

	if !File.IsExist(zipDirPath+"/unzipDir"+dir1) || !File.IsDir(zipDirPath+"/unzipDir"+dir1) {
		t.Error("Unzip failed")
		return
	}

	if !File.IsExist(zipDirPath+"/unzipDir"+dir2) || !File.IsDir(zipDirPath+"/unzipDir"+dir2) {
		t.Error("Unzip failed")
		return
	}

	if !File.IsExist(zipDirPath+"/unzipDir"+dir2+"/file") || File.IsDir(zipDirPath+"/unzipDir"+dir2+"/file") {
		t.Error("Unzip failed")
		return
	}
}

func TestMain(m *testing.M) {
	retCode := m.Run()

	// clean test data
	os.RemoveAll(zipDirPath + ".zip")
	os.RemoveAll(zipDirPath)

	os.RemoveAll(tarDirPath + ".tar")
	os.RemoveAll(tarDirPath)
	os.RemoveAll(untarDirPath)

	os.Exit(retCode)
}
