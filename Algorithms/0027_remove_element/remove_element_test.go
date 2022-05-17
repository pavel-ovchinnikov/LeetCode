package remove_element_0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputData struct {
	nums []int
	val  int
}

type outputData struct {
	result int
}

type TestCase struct {
	in  inputData
	out outputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	testcase := []TestCase{
		{
			inputData{[]int{1, 1, 2}, 2},
			outputData{2},
		},
		{
			inputData{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			outputData{5},
		},
	}

	for _, q := range testcase {
		in, out := q.in, q.out
		ast.Equal(out.result, removeElement(in.nums, in.val), "Input data:%v", in)
	}
}
