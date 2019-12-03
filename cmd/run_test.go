package cmd

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	tt := []struct {
		input  int
		output string
	}{
		{input: 0, output: "result: 1"},
		{input: 1, output: "result: 1"},
		{input: 2, output: "result: 2"},
		{input: 5, output: "result: 120"},
		{input: 10, output: "result: 3628800"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintln("factorial of", tc.input), func(t *testing.T) {
			var buf bytes.Buffer
			factorial(tc.input, func(i int) { buf.WriteString("result: " + strconv.Itoa(i)) })
			assert.Equal(t, tc.output, buf.String())
		})
	}
}
