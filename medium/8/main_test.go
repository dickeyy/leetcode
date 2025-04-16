package main

import "testing"

func Test1(t *testing.T) {
	got := myAtoi("42")
	want := 42

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := myAtoi("   -42")
	want := -42

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := myAtoi("1337c0d3")
	want := 1337

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	got := myAtoi("-91283472332")
	want := -2147483648

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test5(t *testing.T) {
	got := myAtoi("-5-")
	want := -5

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
