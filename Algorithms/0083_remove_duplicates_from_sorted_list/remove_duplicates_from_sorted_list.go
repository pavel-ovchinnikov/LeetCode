package remove_duplicates_from_sorted_list_0083

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	currNode := head
	nextNode := currNode
	for nextNode != nil {
		if currNode.Val != nextNode.Val {
			currNode.Next = nextNode
			currNode = nextNode
		}
		nextNode = nextNode.Next
	}
	currNode.Next = nil
	return head
}
