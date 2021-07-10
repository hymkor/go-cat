package cat

import (
	"fmt"
	"io"
	"os"
)

type eachReader struct {
	fname  string
	reader io.ReadCloser
}

type EachFileError struct {
	fname string
	err   error
}

func (e EachFileError) Error() string {
	return fmt.Sprintf("%s: %s", e.fname, e.err.Error())
}

func (e *EachFileError) Unwrap() error {
	return e.err
}

func (L *eachReader) Read(p []byte) (int, error) {
	if L.reader == nil {
		var err error
		L.reader, err = os.Open(L.fname)
		if err != nil {
			return 0, &EachFileError{fname: L.fname, err: err}
		}
	}
	n, err := L.reader.Read(p)
	if err == nil {
		return n, err
	}
	L.reader.Close()
	if err == io.EOF {
		return n, err
	}
	return n, &EachFileError{fname: L.fname, err: err}
}

func NewReader(files ...string) (io.Reader, error) {
	if len(files) == 0 {
		return os.Stdin, nil
	}
	readers := make([]io.Reader, len(files))
	for i, fname := range files {
		readers[i] = &eachReader{fname: fname}
	}
	return io.MultiReader(readers...), nil
}
