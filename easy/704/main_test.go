package main

import "testing"

func Test1(t *testing.T) {
	got := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	want := 4

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	want := -1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
