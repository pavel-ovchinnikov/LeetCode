package two_sum_0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	one []int
	two int
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
			InputData{[]int{3, 2, 4}, 6},
			OutputData{[]int{1, 2}},
		},
		{
			InputData{[]int{3, 2, 4}, 8},
			OutputData{nil},
		},
	}

	for _, q := range testCase {
		out, in := q.out, q.in
		ast.Equal(out.result, twoSum(in.one, in.two), "Enter:%v", in)
	}
}
