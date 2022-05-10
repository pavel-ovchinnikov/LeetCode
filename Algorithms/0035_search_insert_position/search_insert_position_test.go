package search_insert_position_0035

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	nums   []int
	target int
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
			InputData{[]int{1, 3, 5, 6}, 5},
			OutputData{2},
		},
		{
			InputData{[]int{1, 3, 5, 6}, 2},
			OutputData{1},
		},
		{
			InputData{[]int{2, 3, 5, 6}, 1},
			OutputData{0},
		},
		{
			InputData{[]int{1, 3, 5, 6}, 7},
			OutputData{4},
		},
		{
			InputData{[]int{1, 3, 5, 6, 7}, 8},
			OutputData{5},
		},
		{
			InputData{[]int{1, 3}, 2},
			OutputData{1},
		},
		{
			InputData{[]int{1, 3}, 0},
			OutputData{0},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, searchInsert(in.nums, in.target), "Input data:%v", in)
	}
}
