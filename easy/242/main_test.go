package main

import "testing"

func Test1(t *testing.T) {
	got := isAnagram("anagram", "nagaram")
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
