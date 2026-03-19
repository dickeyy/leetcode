package main

import "testing"

func Test1(t *testing.T) {
	got := isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}})
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}})
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
