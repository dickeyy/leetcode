package main

import "testing"

func Test1(t *testing.T) {
	got := bitwiseComplement(5)
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := bitwiseComplement(7)
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := bitwiseComplement(10)
	want := 5

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	got := bitwiseComplement(0)
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
