package longest_palindromic_substring_0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value string
}

type OutputData struct {
	result string
}

type TestCase struct {
	in  InputData
	out OutputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	testCase := []TestCase{
		{
			InputData{"babad"},
			OutputData{"bab"},
		},
		{
			InputData{"cbbd"},
			OutputData{"bb"},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out
		ast.Equal(out.result, longestPalindrome(in.value), "Input data:%v", in)
	}
}
