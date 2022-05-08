package palindrome_number_0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputData struct {
	value int
}

type outputData struct {
	result bool
}

type question struct {
	in  inputData
	out outputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			in:  inputData{value: -1000},
			out: outputData{result: false},
		},
		{
			in:  inputData{value: 0},
			out: outputData{result: true},
		},
		{
			in:  inputData{value: 10},
			out: outputData{result: false},
		},
		{
			in:  inputData{value: 120321},
			out: outputData{result: false},
		},
		{
			in:  inputData{value: 12321},
			out: outputData{result: true},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, isPalindrome(in.value), "Input data:%v", in)
	}
}
