package main

import "testing"

func Test1(t *testing.T) {
	got := finalValueAfterOperations([]string{"X++", "X++", "X--", "++X", "--X"})
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
