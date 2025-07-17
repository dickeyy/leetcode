package main

import "testing"

func Test1(t *testing.T) {
	got := mySqrt(4)
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := mySqrt(8)
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
