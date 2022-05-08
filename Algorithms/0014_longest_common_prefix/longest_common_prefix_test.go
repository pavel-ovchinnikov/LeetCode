package longest_common_prefix_0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputData struct {
	value []string
}

type outputData struct {
	result string
}

type question struct {
	in  inputData
	out outputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			inputData{[]string{"qweasd", "qwezxc", "qwgh"}},
			outputData{"qw"},
		},
		{
			inputData{[]string{"qwerty", "qwefgh"}},
			outputData{"qwe"},
		},
		{
			inputData{[]string{"aa", "bb", "cc"}},
			outputData{""},
		},
		{
			inputData{[]string{}},
			outputData{""},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, longestCommonPrefix(in.value), "Input data:%v", in)
	}
}
