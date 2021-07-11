package cat_test

import (
	"bytes"
	"errors"
	//"fmt"
	"io"
	"testing"

	"github.com/zetamatta/go-cat"
)

func TestNewReader(t *testing.T) {
	r, err := cat.NewReader("testfile1.txt", "testfile2.txt", "testfile3.txt")
	if err != nil {
		t.Fatalf("NewReader: %s", err.Error())
		return
	}

	bin, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll: %s", err.Error())
		return
	}

	if !bytes.Equal(bin, []byte("1234\r\n5678\r\n9ABC\r\n")) {
		t.Fatalf("not equal: actual: `%s`", string(bin))
		return
	}

	r, err = cat.NewReader("(notfound)")
	if err != nil {
		t.Fatalf("NewReader: %s", err.Error())
		return
	}
	bin, err = io.ReadAll(r)
	if err == nil {
		t.Fatalf("ReadAll: errors should be found, but not found")
		return
	}
	var eachFileError *cat.EachFileError
	if !errors.As(err, &eachFileError) {
		t.Fatalf("NewReader: %s", err.Error())
		return
	}

	r, err = cat.NewReader(".")
	if err != nil {
		t.Fatalf("NewReader: %s", err.Error())
		return
	}
	bin, err = io.ReadAll(r)
	if err == nil {
		t.Fatalf("ReadAll: errors should be found, but not found")
		return
	}
	//fmt.Printf("%s (%#v)\n", err.Error(), err)
}
