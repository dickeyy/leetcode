package main

import "testing"

func Test1(t *testing.T) {
	got := singleNumber([]int{2, 2, 1})
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := singleNumber([]int{4, 1, 2, 1, 2})
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := singleNumber([]int{1})
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
