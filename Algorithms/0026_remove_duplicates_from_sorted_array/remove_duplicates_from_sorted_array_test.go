package remove_duplicates_from_sorted_array_0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputData struct {
	value []int
}

type outputData struct {
	result int
}

type question struct {
	in  inputData
	out outputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			inputData{[]int{1, 1, 2}},
			outputData{2},
		},
		{
			inputData{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			outputData{5},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, removeDuplicates(in.value), "Input data:%v", in)
	}
}
