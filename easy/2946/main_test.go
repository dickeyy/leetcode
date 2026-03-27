package main

import "testing"

func Test1(t *testing.T) {
	got := areSimilar([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1)
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
