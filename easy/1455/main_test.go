package main

import "testing"

func Test1(t *testing.T) {
	got := isPrefixOfWord("I love eating chips", "chips")
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := isPrefixOfWord("I love eating chips", "love")
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
