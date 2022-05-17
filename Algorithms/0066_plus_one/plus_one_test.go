package plus_one_0066

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	digits []int
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
			InputData{[]int{1, 2, 3}},
			OutputData{[]int{1, 2, 4}},
		},
		{
			InputData{[]int{1, 2, 9}},
			OutputData{[]int{1, 3, 0}},
		},
		{
			InputData{[]int{9, 9}},
			OutputData{[]int{1, 0, 0}},
		},
		{
			InputData{[]int{0}},
			OutputData{[]int{1}},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out
		ast.Equal(out.result, plusOne(in.digits), "Input data:%v", in)
	}
}
