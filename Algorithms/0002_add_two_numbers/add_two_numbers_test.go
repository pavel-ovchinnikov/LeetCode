package add_two_numbers_0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value1 []int
	value2 []int
}

type OutputData struct {
	result []int
}

type TestCase struct {
	in  InputData
	out OutputData
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	testCase := []TestCase{
		{
			InputData{[]int{2, 4, 3}, []int{5, 6, 4}},
			OutputData{[]int{7, 0, 8}},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out
		l1 := makeListNode(in.value1)
		l2 := makeListNode(in.value2)
		result := makeListNode(out.result)
		ast.Equal(result, addTwoNumbers(l1, l2), "Input data:%v", in)
	}
}
