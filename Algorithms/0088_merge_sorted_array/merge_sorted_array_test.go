package merge_sorted_array_0088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

type OutputData struct {
	result []int
}

type TastCase struct {
	in  InputData
	out OutputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	tastCase := []TastCase{
		// {
		// 	InputData{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		// 	OutputData{[]int{1, 2, 2, 3, 5, 6}},
		// },
		{
			InputData{[]int{0}, 0, []int{1}, 1},
			OutputData{[]int{1}},
		},
		{
			InputData{[]int{0, 0}, 0, []int{1, 2}, 2},
			OutputData{[]int{1, 2}},
		},
		// {
		// 	InputData{[]int{1}, 1, []int{}, 0},
		// 	OutputData{[]int{1}},
		// },
	}

	for _, test := range tastCase {
		in, out := test.in, test.out

		merge(in.nums1, in.m, in.nums2, in.n)

		ast.Equal(out.result, in.nums1, "Input data:%v", in)
	}
}
