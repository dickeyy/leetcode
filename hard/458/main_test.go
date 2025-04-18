package main

import "testing"

func Test1(t *testing.T) {
	got := poorPigs(90, 2, 7)
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := poorPigs(1, 1, 1)
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
