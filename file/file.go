package file

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gongcf/gtool/logger"
	"github.com/gongcf/gtool/rand"
	"github.com/gongcf/gtool/str"
)

type GFile struct{}

// var logger logger.Logger

// RemoveEmptyDirs removes all empty dirs under the specified dir path.
func (gf *GFile) RemoveEmptyDirs(dir string, excludes ...string) (err error) {
	_, err = gf.removeEmptyDirs(dir, excludes...)
	return
}

func (gf *GFile) removeEmptyDirs(dir string, excludes ...string) (removed bool, err error) {
	// Credit to: https://github.com/InfuseAI/ArtiVC/blob/main/internal/core/utils.go
	// LICENSE Apache License 2.0 https://github.com/InfuseAI/ArtiVC/blob/main/LICENSE
	Str := str.Gstr{}
	dirName := filepath.Base(dir)
	if Str.Contains(dirName, excludes) {
		return
	}

	var hasEntries bool
	entires, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, entry := range entires {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			removed, err = gf.removeEmptyDirs(subdir, excludes...)
			if err != nil {
				return false, err
			}
			if !removed {
				hasEntries = true
			}
		} else {
			hasEntries = true
		}
	}

	if !hasEntries && !Str.Contains(dirName, excludes) {
		err = os.Remove(dir)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (*GFile) IsValidFilename(name string) bool {
	reserved := []string{"\\", "/", ":", "*", "?", "\"", "'", "<", ">", "|"}
	for _, r := range reserved {
		if strings.Contains(name, r) {
			return false
		}
	}
	return true
}

// WriteFileSaferByReader writes the data to a temp file and atomically move if everything else succeeds.
func (*GFile) WriteFileSaferByReader(writePath string, reader io.Reader, perm os.FileMode) (err error) {
	Rand := rand.GRand{}
	dir, name := filepath.Split(writePath)
	tmp := filepath.Join(dir, name+Rand.String(7)+".tmp")
	f, err := os.OpenFile(tmp, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if nil != err {
		return
	}

	if _, err = io.Copy(f, reader); nil != err {
		return
	}

	if err = f.Sync(); nil != err {
		return
	}

	if err = f.Close(); nil != err {
		return
	}

	if err = os.Chmod(f.Name(), perm); nil != err {
		return
	}

	for i := 0; i < 3; i++ {
		err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
		if nil == err {
			os.Remove(f.Name())
			return
		}

		if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	return
}

// WriteFileSaferWithoutChangeTime writes the data to a temp file and atomically move if everything else succeeds, do not change the file modification time.
func (*GFile) WriteFileSaferWithoutChangeTime(writePath string, data []byte, perm os.FileMode) (err error) {
	// credits: https://github.com/vitessio/vitess/blob/master/go/ioutil2/ioutil.go
	Rand := rand.GRand{}
	dir, name := filepath.Split(writePath)
	tmp := filepath.Join(dir, name+Rand.String(7)+".tmp")
	f, err := os.OpenFile(tmp, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if nil != err {
		return
	}

	if _, err = f.Write(data); nil != err {
		return
	}

	if err = f.Sync(); nil != err {
		return
	}

	if err = f.Close(); nil != err {
		return
	}

	if err = os.Chmod(f.Name(), perm); nil != err {
		return
	}

	// 保持文件修改时间不变
	info, _ := os.Stat(writePath)
	if nil != info {
		t := info.ModTime()
		if err = os.Chtimes(f.Name(), t, t); nil != err {
			return
		}
	}

	for i := 0; i < 3; i++ {
		err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
		if nil == err {
			os.Remove(f.Name())
			return
		}

		if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	return
}

// WriteFileSafer writes the data to a temp file and atomically move if everything else succeeds.
func (*GFile) WriteFileSafer(writePath string, data []byte, perm os.FileMode) (err error) {
	// credits: https://github.com/vitessio/vitess/blob/master/go/ioutil2/ioutil.go
	Rand := rand.GRand{}
	dir, name := filepath.Split(writePath)
	tmp := filepath.Join(dir, name+Rand.String(7)+".tmp")
	f, err := os.OpenFile(tmp, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if nil != err {
		return
	}

	if _, err = f.Write(data); nil != err {
		return
	}

	if err = f.Sync(); nil != err {
		return
	}

	if err = f.Close(); nil != err {
		return
	}

	if err = os.Chmod(f.Name(), perm); nil != err {
		return
	}

	for i := 0; i < 3; i++ {
		err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
		if nil == err {
			os.Remove(f.Name())
			return
		}

		if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	return
}

// GetFileSize get the length in bytes of file of the specified path.
func (*GFile) GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		logger.Glogger.Error(err)

		return -1
	}

	return fi.Size()
}

// IsExist determines whether the file spcified by the given path is exists.
func (*GFile) IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

// IsBinary determines whether the specified content is a binary file content.
func (*GFile) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}

	return false
}

// IsImg determines whether the specified extension is a image.
func (*GFile) IsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// IsDir determines whether the specified path is a directory.
func (*GFile) IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		logger.Glogger.Warnf("determines whether [%s] is a directory failed: [%v]", path, err)
		return false
	}
	return fio.IsDir()
}

// Copy copies the source to the dest.
// Keep the dest access/mod time as the same as the source.
func (gf *GFile) Copy(source, dest string) (err error) {
	if !gf.IsExist(source) {
		return os.ErrNotExist
	}

	if gf.IsDir(source) {
		return gf.CopyDir(source, dest)
	}
	return gf.CopyFile(source, dest)
}

// CopyNewtimes copies the source to the dest.
// Do not keep the dest access/mod time as the same as the source.
func (gf *GFile) CopyNewtimes(source, dest string) (err error) {
	if !gf.IsExist(source) {
		return os.ErrNotExist
	}

	if gf.IsDir(source) {
		return gf.CopyDirNewtimes(source, dest)
	}
	return gf.CopyFileNewtimes(source, dest)
}

// CopyFile copies the source file to the dest file.
// Keep the dest access/mod time as the same as the source.
func (gf *GFile) CopyFile(source, dest string) (err error) {
	return gf.copyFile(source, dest, true)
}

// CopyFileNewtimes copies the source file to the dest file.
// Do not keep the dest access/mod time as the same as the source.
func (gf *GFile) CopyFileNewtimes(source, dest string) (err error) {
	return gf.copyFile(source, dest, false)
}

func (*GFile) copyFile(source, dest string, chtimes bool) (err error) {
	if err = os.MkdirAll(filepath.Dir(dest), 0755); nil != err {
		return
	}

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()

	sourceinfo, err := os.Stat(source)
	if nil != err {
		return
	}

	if err = os.Chmod(dest, sourceinfo.Mode()); nil != err {
		return
	}

	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourcefile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if nil != err {
		return
	}

	if chtimes {
		if err = os.Chtimes(dest, sourceinfo.ModTime(), sourceinfo.ModTime()); nil != err {
			return
		}
	}
	return
}

// CopyDir copies the source directory to the dest directory.
// Keep the dest access/mod time as the same as the source.
func (gf *GFile) CopyDir(source, dest string) (err error) {
	return gf.copyDir(source, dest, true)
}

// CopyDirNewtimes copies the source directory to the dest directory.
// Do not keep the dest access/mod time as the same as the source.
func (gf *GFile) CopyDirNewtimes(source, dest string) (err error) {
	return gf.copyDir(source, dest, false)
}

func (gf *GFile) copyDir(source, dest string, chtimes bool) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	dirs, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, f := range dirs {
		srcFilePath := filepath.Join(source, f.Name())
		destFilePath := filepath.Join(dest, f.Name())

		if f.IsDir() {
			err = gf.copyDir(srcFilePath, destFilePath, chtimes)
			if err != nil {
				logger.Glogger.Error(err)
				return
			}
		} else {
			err = gf.copyFile(srcFilePath, destFilePath, chtimes)
			if err != nil {
				logger.Glogger.Error(err)
				return
			}
		}
	}

	if chtimes {
		if err = os.Chtimes(dest, sourceinfo.ModTime(), sourceinfo.ModTime()); nil != err {
			return
		}
	}
	return nil
}

func (*GFile) IsHidden(path string) bool {
	path = filepath.Base(path)
	if 1 > len(path) {
		return false
	}
	return "." == path[:1]
}
