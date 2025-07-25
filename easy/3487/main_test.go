package main

import "testing"

func Test1(t *testing.T) {
	got := maxSum([]int{1, 2, 3, 4, 5})
	want := 15

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := maxSum([]int{-100})
	want := -100

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := maxSum([]int{-20, 20})
	want := 20

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	got := maxSum([]int{2, -10, 6})
	want := 8

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
