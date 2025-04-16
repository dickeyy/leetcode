package main

import "testing"

func Test1(t *testing.T) {
	got := subsetXORSum([]int{1, 3})
	want := 6

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := subsetXORSum([]int{5, 1, 6})
	want := 28

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := subsetXORSum([]int{3, 4, 5, 6, 7, 8})
	want := 480

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
