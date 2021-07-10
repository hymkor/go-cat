//+build ignore

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/zetamatta/go-cat"
)

func main() {
	r, err := cat.NewReader("testfile1.txt", "testfile2.txt", "testfile3.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	bin, err := io.ReadAll(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	io.WriteString(os.Stdout, string(bin))
}
