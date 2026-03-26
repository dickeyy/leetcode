package main

import "testing"

func Test1(t *testing.T) {
	got := containsDuplicate([]int{1, 2, 3, 1})
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
