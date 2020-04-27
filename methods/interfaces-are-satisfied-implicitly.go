package main

import (
	"fmt"
	"os"
)

// Reader interface
type Reader interface{
	Read(b []byte)(n int, err error)
}

// Writer interface
type Writer interface{
	Write(b []byte)(n int, err error)
}

// ReadWriter interface
type ReadWriter interface{
	Reader
	Writer
}

func main(){
	var w Writer

	// os.Stdout 实现了 Writer
	w = os.Stdout
	fmt.Fprintf(w, "Hello, Writer\n")
}