package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (t rot13Reader) Read(b []byte) (int, error) {
	n, err := t.r.Read(b)
	for i, c := range b {
		b[i] = convertRot13(c)
	}
	return n, err
}

const UpperCase byte = 65
const LowerCase byte = 97

func convertRot13(b byte) byte {
	if b < LowerCase {
		b %= UpperCase
	} else {
		b %= LowerCase
	}
	return b % 13
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
