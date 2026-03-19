package main

/*
 *   Stats:
 *   Runtime: 2ms (beats 78.54% of Go solutions)
 *   Memory: 10.62mb (beats 69.72% of Go solutions)
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var rev *ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		temp := slow.Next
		slow.Next = rev
		rev = slow
		slow = temp
	}

	if fast != nil {
		slow = slow.Next
	}

	for rev != nil {
		if rev.Val != slow.Val {
			return false
		}
		rev = rev.Next
		slow = slow.Next
	}

	return true
}
