package longest_substring_without_repeating_characters_0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value string
}

type OutputData struct {
	result int
}

type TestCase struct {
	in  InputData
	out OutputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	testCase := []TestCase{
		{
			InputData{"abcabcbb"},
			OutputData{3},
		},
		{
			InputData{"aaaaaaa"},
			OutputData{1},
		},
		{
			InputData{"pwwkew"},
			OutputData{3},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out
		ast.Equal(out.result, lengthOfLongestSubstring(in.value), "Input data:%v", in)
	}
}
