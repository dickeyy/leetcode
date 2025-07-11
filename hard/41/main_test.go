package main

import "testing"

func Test1(t *testing.T) {
	got := firstMissingPositive([]int{1, 2, 0})
	want := 3

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := firstMissingPositive([]int{3, 4, -1, 1})
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := firstMissingPositive([]int{7, 8, 9, 11, 12})
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	got := firstMissingPositive([]int{1})
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test5(t *testing.T) {
	got := firstMissingPositive([]int{1, 1})
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
