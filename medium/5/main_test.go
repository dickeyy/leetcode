package main

import "testing"

func Test1(t *testing.T) {
	got := longestPalindrome("babad")
	want := "bab"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := longestPalindrome("cbbd")
	want := "bb"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
