package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	head := &ListNode{}
	dummy := head
	for l1 != nil || l2 != nil || carry > 0 {
		cnt := carry
		if l1 != nil {
			cnt += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			cnt += l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{
			cnt % 10,
			nil,
		}
		head = head.Next
		carry = cnt / 10
	}
	return dummy.Next
}
