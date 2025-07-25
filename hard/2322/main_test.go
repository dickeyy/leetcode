package main

import (
	"testing"
)

func Test1(t *testing.T) {
	got := minimumScore([]int{1, 5, 5, 4, 11}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}})
	want := 9

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := minimumScore([]int{5, 5, 2, 4, 4, 2}, [][]int{{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3}})
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
