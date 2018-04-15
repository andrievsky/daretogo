package main

import (
	"fmt"

	"github.com/andrievsky/daretogo/challenges/helpers"
)

const count = 6

func main() {
	var reader = helpers.NewReader()
	var s = make([][]int, count, count)
	for i := 0; i < count; i++ {
		s[i] = reader.ReadIntSlice(count)
	}

	fmt.Print(s)
}
