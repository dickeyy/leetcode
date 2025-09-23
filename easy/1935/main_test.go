package main

import "testing"

func Test1(t *testing.T) {
	got := canBeTypedWords("hello world", "ad")
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
