package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	tt := []struct {
		input  int
		output int
	}{
		{input: 0, output: 1},
		{input: 1, output: 1},
		{input: 2, output: 2},
		{input: 5, output: 120},
		{input: 10, output: 3628800},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintln("factorial of", tc.input), func(t *testing.T) {
			factorial(tc.input, func(i int) {
				assert.Equal(t, tc.output, i)
			})
		})
	}
}
