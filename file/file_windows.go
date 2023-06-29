package file

// IsHidden checks whether the file specified by the given path is hidden.
func (*GFile) IsHidden(path string) bool {
	pointer, err := syscall.UTF16PtrFromString(path)
	if nil != err {
		logger.Errorf("Checks file [%s] is hidden failed: [%s]", path, err)

		return false
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if nil != err {
		logger.Errorf("Checks file [%s] is hidden failed: [%s]", path, err)

		return false
	}

	return 0 != attributes&syscall.FILE_ATTRIBUTE_HIDDEN
}
