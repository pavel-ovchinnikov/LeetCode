package maximum_subarray_0053

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	nums []int
}

type OutputData struct {
	result int
}

type TastCase struct {
	in  InputData
	out OutputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []TastCase{
		{
			InputData{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			OutputData{6},
		},
		{
			InputData{[]int{5, 4, -1, 7, 8}},
			OutputData{23},
		},
		{
			InputData{[]int{1}},
			OutputData{1},
		},
		{
			InputData{[]int{-1}},
			OutputData{-1},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, maxSubArray(in.nums), "Input data:%v", in)
	}
}
