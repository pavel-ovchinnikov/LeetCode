package valid_parentheses_0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputData struct {
	value string
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
			inputData{"(()(()()))"},
			outputData{true},
		},
		{
			inputData{"()"},
			outputData{true},
		},
		{
			inputData{")()"},
			outputData{false},
		},
		{
			inputData{"()()"},
			outputData{true},
		},
		{
			inputData{"()(()"},
			outputData{false},
		},
		{
			inputData{"()[]{}"},
			outputData{true},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, isValid(in.value), "Input data:%v", in)
	}
}
