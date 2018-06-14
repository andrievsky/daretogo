package main

import (
	"testing"

	"github.com/andrievsky/daretogo/algo"
	"github.com/stretchr/testify/assert"
)

func TestCheckForLetterAndDigit(t *testing.T) {
	for i, c := range algo.LetterAndDigitCases {
		assert.Equal(t, algo.LetterAndDigitResults[i], algo.CheckForLetterAndDigit(c))
	}
}

func BenchmarkCheckForLetterAndDigit(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		for _, c := range algo.LetterAndDigitCases {
			algo.CheckForLetterAndDigit(c)
		}
	}
}

/*
func TestCheckForLetterAndDigitRegex(t *testing.T) {
	for i, c := range algo.LetterAndDigitCases {
		assert.Equal(t, algo.LetterAndDigitResults[i], algo.CheckForLetterAndDigitWithRegex(c))
	}
}
*/
