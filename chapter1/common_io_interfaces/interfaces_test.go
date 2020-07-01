package common_io_interfaces

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	io := ioout{
		bytes.NewReader([]byte("example")),
		&bytes.Buffer{},
	}

	fmt.Print("stdout on Copy = ")

	if err := Copy(io.in, io.out); err != nil {
		t.Error(err)
	}
	fmt.Println("out bytes buffer =", io.out.String())

}
