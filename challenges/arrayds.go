package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var s = read()
	reverse(s)

	fmt.Print(format(s))
}

func format(s []int) string {
	var buffer bytes.Buffer

	for i := 0; i < len(s)-1; i++ {
		buffer.WriteString(strconv.Itoa(s[i]))
		buffer.WriteString(" ")
	}
	buffer.WriteString(strconv.Itoa(s[len(s)-1]))
	return buffer.String()
}

func reverse(s []int) {
	var j = len(s) - 1
	for i := 0; i <= len(s)/2-1; i++ {
		s[i], s[j-i] = s[j-i], s[i]
	}
}

func read() []int {
	var scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)
	var count = nextInt(scanner)
	var source = make([]int, count, count)
	for i := 0; i < count; i++ {
		source[i] = nextInt(scanner)
	}
	return source
}

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatalln("Reading input:", err, os.Stderr)
	}
	value, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	return value
}
