package main

import "testing"

func Test1(t *testing.T) {
	got := minOperations("0100")
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := minOperations("10")
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := minOperations("1111")
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
