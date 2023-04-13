package customerr

import (
	"fmt"
	"os"
)

type FileError struct {
	code int
	msg  string
}

func (f *FileError) Error() string {
	return fmt.Sprintf("code = %d,msg = %s", f.code, f.msg)
}

var _ error = (*FileError)(nil)

func OpenFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return &FileError{
			code: 201,
			msg:  err.Error(),
		}
	}
	defer file.Close()
	return nil
}
