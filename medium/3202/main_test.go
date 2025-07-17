package main

import (
	"testing"
)

func Test1(t *testing.T) {
	got := maximumLength([]int{1, 2, 3, 4, 5}, 2)
	want := 5

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := maximumLength([]int{1, 8, 3, 8, 3}, 4)
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
