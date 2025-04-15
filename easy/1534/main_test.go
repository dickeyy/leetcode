package main

import "testing"

func Test1(t *testing.T) {
	arr := []int{3, 0, 1, 1, 9, 7}
	a := 7
	b := 2
	c := 3
	got := countGoodTriplets(arr, a, b, c)
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	arr := []int{1, 1, 2, 2, 3}
	a := 0
	b := 0
	c := 1
	got := countGoodTriplets(arr, a, b, c)
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
