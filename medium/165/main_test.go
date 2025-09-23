package main

import "testing"

func Test1(t *testing.T) {
	got := compareVersion("1.2", "1.10")
	want := -1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := compareVersion("1.0", "1.0.0.0")
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := compareVersion("1.0.1", "1")
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	got := compareVersion("0.1", "1.1")
	want := -1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
