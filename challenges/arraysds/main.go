package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/andrievsky/daretogo/challenges/helpers"
)

func main() {
	var reader = helpers.NewReader()
	var count = reader.ReadInt()
	var s = reader.ReadIntSlice(count)
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
