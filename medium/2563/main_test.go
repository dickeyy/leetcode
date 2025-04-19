package main

import "testing"

func Test1(t *testing.T) {
	got := countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6)
	want := int64(6)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11)
	want := int64(1)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
