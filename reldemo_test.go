package reldemo

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestDumpBytes(t *testing.T) {
	output := "var foobar = []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}\n"
	dump, err := DumpBytes(bytes.NewBufferString("hello world"), "foobar")
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, dump)
	result := string(buf.Bytes())
	if result != output {
		t.Errorf("DumpBytes()  = \"%v\", want \"%v\".", result, output)
	}
}

func ExampleDumpBytes() {
	dump, err := DumpBytes(bytes.NewBufferString("hello world"), "foobar")
	if err != nil {
		return
	}
	io.Copy(os.Stdout, dump)
	// Output:
	// var foobar = []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}
}
