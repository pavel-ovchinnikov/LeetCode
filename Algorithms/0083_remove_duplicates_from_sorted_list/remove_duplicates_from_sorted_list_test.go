package remove_duplicates_from_sorted_list_0083

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value []int
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
		{
			InputData{[]int{1, 1, 2}},
			OutputData{[]int{1, 2}},
		},
		{
			InputData{[]int{1, 1}},
			OutputData{[]int{1}},
		},
		{
			InputData{[]int{1, 1, 1, 3}},
			OutputData{[]int{1, 3}},
		},
	}

	for _, test := range tastCase {
		in, out := test.in, test.out
		head := sliceTolist(in.value)
		ast.Equal(out.result, listToSlice(deleteDuplicates(head)), "Input data:%v", in)
	}
}

func listToSlice(head *ListNode) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func sliceTolist(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}
