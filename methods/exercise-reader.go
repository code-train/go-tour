package main

import (
	"golang.org/x/tour/reader"
)

// MyReader struct
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main(){
	reader.Validate(MyReader{})
}