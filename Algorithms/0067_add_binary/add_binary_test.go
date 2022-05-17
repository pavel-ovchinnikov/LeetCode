package add_binary_0067

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value1 string
	value2 string
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
			InputData{"11", "11"},
			OutputData{"110"},
		},
		{
			InputData{"11", "1"},
			OutputData{"100"},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out
		ast.Equal(out.result, addBinary(in.value1, in.value2), "Input data:%v", in)
	}
}
