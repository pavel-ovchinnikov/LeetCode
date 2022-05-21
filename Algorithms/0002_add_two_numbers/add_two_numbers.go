package add_two_numbers_0002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := &ListNode{}
	currNode := headNode

	sum := 0

	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		currNode.Next = &ListNode{Val: sum % 10}
		currNode = currNode.Next

		sum /= 10
	}
	return headNode.Next
}
