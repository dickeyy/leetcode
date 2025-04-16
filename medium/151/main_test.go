package main

import "testing"

func Test1(t *testing.T) {
	got := reverseWords("the sky is blue")
	want := "blue is sky the"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := reverseWords("  hello world  ")
	want := "world hello"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := reverseWords("a good   example")
	want := "example good a"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
