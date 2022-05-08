package roman_to_integer_0013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputData struct {
	value string
}

type outputData struct {
	result int
}

type question struct {
	in  inputData
	out outputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			inputData{"I"},
			outputData{1},
		},
		{
			inputData{"IV"},
			outputData{4},
		},
		{
			inputData{"V"},
			outputData{5},
		},
		{
			inputData{"VII"},
			outputData{7},
		},
		{
			inputData{"MDCCCLXXXVIII"},
			outputData{1888},
		},
		{
			inputData{"MMMCMXCIX"},
			outputData{3999},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, romanToInt(in.value), "Input data:%v", in)
	}
}
