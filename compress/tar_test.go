package compress

import (
	"path/filepath"
	"testing"
)

var tarDirPath = testdataDir + "_tar"
var untarDirPath = testdataDir + "_untar"

var Tar GCompress

func TestTar(t *testing.T) {
	if err := Tar.Tar(testdataDir, tarDirPath); nil != err {
		t.Error(err)
		return
	}
}

func TestUntar(t *testing.T) {
	err := Tar.Untar(filepath.Join(tarDirPath, testdataDir+".tar"), untarDirPath)
	if nil != err {
		t.Error(err)
		return
	}
}
