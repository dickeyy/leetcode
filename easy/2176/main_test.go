package main

import "testing"

func Test1(t *testing.T) {
	got := countPairs([]int{3, 1, 2, 2, 2, 1, 3}, 2)
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := countPairs([]int{1, 2, 3, 4}, 1)
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
