package main

import "testing"

func Test1(t *testing.T) {
	got := minFlips("111000")
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := minFlips("010")
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := minFlips("1110")
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	got := minFlips("01001001101")
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
