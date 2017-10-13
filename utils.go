package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(input ...int) *ListNode {
	head := &ListNode{}
	tmp := head

	for _, v := range input {
		head.Next = &ListNode{v, nil}
		head = head.Next
	}
	return tmp.Next
}

func (l *ListNode) EqualsTo(o *ListNode) bool {
	return l.String() == o.String()
}

func (l *ListNode) String() string {
	sl := []int{}
	for l != nil {
		sl = append(sl, l.Val)
		l = l.Next
	}

	return fmt.Sprintf("%v", sl)
}
