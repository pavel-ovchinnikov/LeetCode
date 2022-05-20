package single_number_0136

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value []int
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
			InputData{[]int{2, 2, 1}},
			OutputData{1},
		},
		{
			InputData{[]int{4, 1, 2, 1, 2}},
			OutputData{4},
		},
		{
			InputData{[]int{1}},
			OutputData{1},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out

		ast.Equal(out.result, singleNumber(in.value), "Input data:%v", in)
	}
}
