package helpers

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// InputReader read data from stdin
type InputReader struct {
	scanner *bufio.Scanner
}

// NewReader InputReader constructor
func NewReader() *InputReader {
	var reader = new(InputReader)
	reader.scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))
	reader.scanner.Split(bufio.ScanWords)
	return reader
}

// ReadIntSlice read int[] where count is a slice lenght
func (t *InputReader) ReadIntSlice(count int) []int {
	var source = make([]int, count, count)
	for i := 0; i < count; i++ {
		source[i] = t.ReadInt()
	}
	return source
}

// ReadInt read int
func (t *InputReader) ReadInt() int {
	t.scanner.Scan()
	if err := t.scanner.Err(); err != nil {
		log.Fatalln("Reading input:", err, os.Stderr)
	}
	value, err := strconv.Atoi(t.scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	return value
}
