package main

import "testing"

func Test1(t *testing.T) {
	got := isValid("234Adas")
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := isValid("b3")
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := isValid("a3$e")
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	got := isValid("UuE6")
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
