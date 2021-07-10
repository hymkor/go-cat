go-cat
======

cat.NewReader is similar with io.MultiReader,
but it opens the each file on demand and closes when the file is read all.

```go
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
```

```
$ cat testfile1.txt
1234
$ cat testfile2.txt
5678
$ cat testfile3.txt
9ABC
$ go run example.go
1234
5678
9ABC
$
```
