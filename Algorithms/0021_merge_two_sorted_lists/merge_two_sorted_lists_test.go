package merge_two_sorted_lists_0021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputData struct {
	list1 []int
	list2 []int
}

type outputData struct {
	result []int
}

type question struct {
	in  inputData
	out outputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			inputData{[]int{1, 2, 4}, []int{1, 3, 4}},
			outputData{[]int{1, 1, 2, 3, 4, 4}},
		},
		{
			inputData{[]int{1, 2, 4}, []int{0, 3}},
			outputData{[]int{0, 1, 2, 3, 4}},
		},
		{
			inputData{[]int{}, []int{0}},
			outputData{[]int{0}},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, listToSlice(mergeTwoLists(sliceTolist(in.list1), sliceTolist(in.list2))), "Input data:%v", in)
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
