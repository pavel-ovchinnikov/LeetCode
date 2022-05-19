package valid_palindrome_0125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value string
}

type OutputData struct {
	result bool
}

type TestCase struct {
	in  InputData
	out OutputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	testCase := []TestCase{
		{
			InputData{"A man, a plan, a canal: Panama"},
			OutputData{true},
		},
		{
			InputData{"race a car"},
			OutputData{false},
		},
		{
			InputData{"  "},
			OutputData{true},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out

		ast.Equal(out.result, isPalindrome(in.value), "Input data:%v", in)
	}
}
