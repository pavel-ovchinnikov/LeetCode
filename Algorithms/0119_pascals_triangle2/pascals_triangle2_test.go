package pascals_triangle_0119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	numRows int
}

type OutputData struct {
	result []int
}

type TestCase struct {
	in  InputData
	out OutputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	testCase := []TestCase{
		{
			InputData{0},
			OutputData{[]int{1}},
		},
		{
			InputData{1},
			OutputData{[]int{1, 1}},
		},
		{
			InputData{2},
			OutputData{[]int{1, 2, 1}},
		},
		{
			InputData{3},
			OutputData{[]int{1, 3, 3, 1}},
		},
		{
			InputData{4},
			OutputData{[]int{1, 4, 6, 4, 1}},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out

		ast.Equal(out.result, getRow(in.numRows), "Input data:%v", in)
	}
}
