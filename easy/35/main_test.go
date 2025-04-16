package main

import "testing"

func Test1(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	got := searchInsert(nums, target)
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
