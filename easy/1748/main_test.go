package main

import "testing"

func Test1(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	got := sumOfUnique(nums)
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
